package controllers

import (
	"context"
	"go-api/src/dtos"
	"go-api/src/services"
	"go-api/src/utils"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	service *services.UserService
}

func NewUserController(server *gin.Engine, service *services.UserService) {
	controller := &UserController{service: service}

	routes := server.Group("/users")
	{
		routes.POST("", controller.CreateUser)
		routes.POST("/login", controller.LoginUser)
	}
}

// @Tags users
// @Router /users [post]
// @Summary Criar um novo usuário
// @Description Registra um novo usuário na API
// @Accept json
// @Produce json
// @Param user body dtos.CreateUserRequest true "Dados do usuário"
// @Success 201 {object} dtos.Message "Usuário criado"
// @Failure 400 {object} dtos.APIError "Erro de validação"
// @Failure 409 {object} dtos.APIError "Usuário já cadastrado"
func (c *UserController) CreateUser(ginCtx *gin.Context) {
	var req dtos.CreateUserRequest
	if err := utils.ValidateRequestBody(ginCtx, &req); err != nil {
		ginCtx.Error(err)
		return
	}

	ctx, cancel := context.WithTimeout(ginCtx.Request.Context(), 5*time.Second)
	defer cancel()

	if err := c.service.RegisterUser(ctx, req.Email, req.Password); err != nil {
		ginCtx.Error(err)
		return
	}

	ginCtx.JSON(http.StatusCreated, dtos.Message{Message: "Usuário criado com sucesso."})
}

func (c *UserController) LoginUser(ginCtx *gin.Context) {
	var req dtos.LoginUserRequest
	err := utils.ValidateRequestBody(ginCtx, &req)
	if err != nil {
		ginCtx.Error(err)
	}

	ctx, cancel := context.WithTimeout(ginCtx.Request.Context(), 5*time.Second)
	defer cancel()

	token, err := c.service.LoginUser(ctx, req.Email, req.Password)
	if err != nil {
		ginCtx.Error(err)
		return
	}

	ginCtx.JSON(http.StatusOK, dtos.LoginUserResponse{Success: true, Token: token})
}
