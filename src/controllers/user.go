package controllers

import (
	"go-api/src/dtos"
	"go-api/src/repositories"
	"go-api/src/services"
	"go-api/src/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	service *services.UserService
}

func NewUserController(server *gin.Engine, repository *repositories.UserRepository) {
	service := services.NewUserService(repository)
	controller := &UserController{service: service}

	routes := server.Group("/users")
	{
		routes.POST("", controller.CreateUser)

	}
}

// @Tags users
// @Router /users [post]
// @Summary Criar um novo usuário
// @Description Registra um novo usuário na API
// @Accept json
// @Produce json
// @Param user body dtos.User true "Dados do usuário"
// @Success 201 {object} dtos.Message "Usuário criado"
// @Failure 400 {object} dtos.APIError "Erro de validação
func (controller *UserController) CreateUser(ginContext *gin.Context) {
	var userDto dtos.User

	err := utils.ValidateRequestBody(ginContext, &userDto)
	if err != nil {
		ginContext.Error(err)
		return
	}

	err = controller.service.CreateUser(userDto.Email, userDto.Password)
	if err != nil {
		ginContext.Error(err)
		return
	}

	ginContext.JSON(http.StatusCreated, dtos.Message{
		Message: "Usuário criado com sucesso.",
	})
}
