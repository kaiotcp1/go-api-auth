package main

import (
	"context"
	_ "go-api/docs"
	"go-api/src/controllers"
	"go-api/src/repositories"
	"go-api/src/services"
	"go-api/src/utils/middleware"
	"log"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// @title API de Autenticação e Tarefas
// @version 1.0
// @description API para gerenciamento de usuários e tarefas usando GIN e Mongodb
// @host 127.0.0.1:8080
// @BasePath /
func main() {
	if err := godotenv.Load(); err != nil {
		log.Println("arquivo .env não encontrado, usando variáveis do ambiente")
	}

	uri := os.Getenv("MONGO_URI")
	dbName := os.Getenv("MONGO_DB_NAME")
	userCollection := os.Getenv("MONGO_USER_COLLECTION")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
	if err != nil {
		log.Fatalf("erro ao conectar ao MongoDB: %v", err)
	}
	if err = client.Ping(ctx, nil); err != nil {
		log.Fatalf("MongoDB não respondeu ao ping: %v", err)
	}

	db := client.Database(dbName)

	userRepo := repositories.NewUserRepository(db.Collection(userCollection))
	userService := services.NewUserService(userRepo)

	server := gin.Default()
	server.Use(middleware.ErrorMiddlewareHandler())

	controllers.NewUserController(server, userService)

	//@securityDefinitions.apikey BearerAuth
	//@in header
	//@name Authorization
	//@description Value: Bearer abc... (Bearer+space+token)
	server.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, ginSwagger.DefaultModelsExpandDepth(-1)))

	server.Run(":8080")
}
