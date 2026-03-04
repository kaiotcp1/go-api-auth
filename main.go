package main

import (
	"go-api/src/repositories"
	"log"
	"net/http"
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
	taskCollection := os.Getenv("MONGO_TASK_COLLECTION")

	repoUser, errUser := repositories.NewUserRepository(uri, dbName, userCollection)
	repoTask, errTask := repositories.NewTaskRepository(uri, dbName, taskCollection)
	app := gin.Default()

	app.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "Hello, Gin!"})
	})

	app.Run(":8080")
}
