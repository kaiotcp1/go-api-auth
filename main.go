package main

import (
	"context"
	"log"

	_ "go-api/docs"
	"go-api/src/config"
	"go-api/src/config/database"
	"go-api/src/controllers"
	"go-api/src/repositories"
	"go-api/src/services"
	"go-api/src/utils/middleware"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title API de Autenticação e Tarefas
// @version 1.0
// @description API para gerenciamento de usuários e tarefas usando GIN e MongoDB
// @host 127.0.0.1:8080
// @BasePath /
func main() {
	appConfig, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("failed to load config: %v", err)
	}

	mongoClient, err := database.NewMongoClient(appConfig.MongoURI)
	if err != nil {
		log.Fatalf("failed to initialize MongoDB: %v", err)
	}
	defer func() {
		if err := mongoClient.Disconnect(context.Background()); err != nil {
			log.Printf("failed to disconnect MongoDB: %v", err)
		}
	}()

	databaseInstance := mongoClient.Database(appConfig.MongoDBName)
	userCollection := databaseInstance.Collection(appConfig.MongoUserCollection)
	userRepository := repositories.NewUserRepository(userCollection)
	jwtService := services.NewJWTService(appConfig.JWTSecret, appConfig.JWTIssuer)
	userService := services.NewUserService(userRepository, jwtService)

	router := gin.Default()
	router.Use(middleware.ErrorMiddlewareHandler())

	controllers.NewUserController(router, userService)

	//@securityDefinitions.apikey BearerAuth
	//@in header
	//@name Authorization
	//@description Value: Bearer abc... (Bearer+space+token)
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, ginSwagger.DefaultModelsExpandDepth(-1)))

	if err := router.Run(":" + appConfig.HTTPPort); err != nil {
		log.Fatalf("failed to start HTTP server: %v", err)
	}
}
