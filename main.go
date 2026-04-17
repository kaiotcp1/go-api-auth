package main

import (
	"context"
	"log"
	"net/http"
	"time"

	docs "go-api/docs"
	"go-api/src/config"
	"go-api/src/config/database"
	"go-api/src/controllers"
	"go-api/src/dtos"
	"go-api/src/repositories"
	"go-api/src/services"
	"go-api/src/utils/middleware"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title Go Auth API
// @version 1.1.0
// @description Portfolio-ready authentication API built with Gin, MongoDB and JWT.
// @contact.name Kaio
// @contact.url https://github.com/
// @license.name MIT
// @host 127.0.0.1:8080
// @BasePath /
// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
// @description Use the format: Bearer <token>
func main() {
	appConfig, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("failed to load config: %v", err)
	}

	docs.SwaggerInfo.Title = appConfig.AppName
	docs.SwaggerInfo.Version = appConfig.AppVersion
	docs.SwaggerInfo.Description = "Portfolio-ready authentication API built with Gin, MongoDB and JWT."
	docs.SwaggerInfo.Host = "127.0.0.1:" + appConfig.HTTPPort

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

	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())
	router.Use(middleware.HeadersMiddleware(appConfig.AppName, appConfig.AppVersion, middleware.ParseAllowedOrigins(appConfig.AllowedOrigins)))
	router.Use(middleware.ErrorMiddlewareHandler())

	router.GET("/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, dtos.APIInfoResponse{
			Name:        appConfig.AppName,
			Version:     appConfig.AppVersion,
			Description: "JWT authentication API built with Gin and MongoDB.",
			DocsURL:     "/swagger/index.html",
			Endpoints:   []string{"/health", "/api/v1/users", "/api/v1/users/login"},
		})
	})

	router.GET("/health", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, dtos.HealthResponse{
			Status:    "ok",
			Service:   appConfig.AppName,
			Version:   appConfig.AppVersion,
			Timestamp: time.Now().UTC().Format(time.RFC3339),
		})
	})

	api := router.Group("/api/v1")
	controllers.NewUserController(api, userService)

	router.NoRoute(func(ctx *gin.Context) {
		ctx.JSON(http.StatusNotFound, dtos.APIError{StatusCode: http.StatusNotFound, Message: "route not found"})
	})

	router.NoMethod(func(ctx *gin.Context) {
		ctx.JSON(http.StatusMethodNotAllowed, dtos.APIError{StatusCode: http.StatusMethodNotAllowed, Message: "method not allowed"})
	})

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, ginSwagger.DefaultModelsExpandDepth(-1)))

	if err := router.Run(":" + appConfig.HTTPPort); err != nil {
		log.Fatalf("failed to start HTTP server: %v", err)
	}
}
