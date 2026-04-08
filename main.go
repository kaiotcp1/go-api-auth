package main

import (
	_ "go-api/docs"
	"go-api/src/controllers"
	"go-api/src/repositories"
	"go-api/src/utils/middleware"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
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
	// taskCollection := os.Getenv("MONGO_TASK_COLLECTION")

	repoUser, errUser := repositories.NewUserRepository(uri, dbName, userCollection)
	// repoTask, errTask := repositories.NewTaskRepository(uri, dbName, taskCollection)

	if errUser != nil {
		log.Fatalf("Erro ao conectar ao banco de dados: %v", errUser)
	}

	server := gin.Default()
	server.Use(middleware.ErrorMiddlewareHandler())

	controllers.NewUserController(server, repoUser)

	//@securityDefinitions.apikey BearerAuth
	//@in header
	//@name Authorization
	//@description Value: Bearer abc... (Bearer+space+token)
	server.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, ginSwagger.DefaultModelsExpandDepth(-1)))
	// controllers.NewTaskController(server, repoTask)

	server.Run(":8080")
}
