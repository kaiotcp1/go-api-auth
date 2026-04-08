package main

import (
	"go-api/src/controllers"
	"go-api/src/repositories"
	"go-api/src/utils/middleware"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

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
	// controllers.NewTaskController(server, repoTask)

	server.Run(":8080")
}
