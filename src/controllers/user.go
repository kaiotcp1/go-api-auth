package controllers

import (
	"context"
	"net/http"
	"time"

	"go-api/src/dtos"
	"go-api/src/services"
	"go-api/src/utils"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	service *services.UserService
}

func NewUserController(server gin.IRouter, service *services.UserService) {
	controller := &UserController{service: service}

	routes := server.Group("/users")
	{
		routes.POST("", controller.CreateUser)
		routes.POST("/login", controller.LoginUser)
	}
}

// @Tags users
// @Router /api/v1/users [post]
// @Summary Create a new user
// @Description Registers a new user account in the API.
// @Accept json
// @Produce json
// @Param user body dtos.CreateUserRequest true "User payload"
// @Success 201 {object} dtos.Message "User created successfully"
// @Failure 400 {object} dtos.APIError "Validation error"
// @Failure 409 {object} dtos.APIError "User already exists"
// @Failure 500 {object} dtos.APIError "Internal server error"
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

	ginCtx.JSON(http.StatusCreated, dtos.Message{Message: "Usuario criado com sucesso."})
}

// @Tags auth
// @Router /api/v1/users/login [post]
// @Summary Login
// @Description Validates credentials and returns a JWT token.
// @Accept json
// @Produce json
// @Param credentials body dtos.LoginUserRequest true "Login payload"
// @Success 200 {object} dtos.LoginUserResponse "Authenticated successfully"
// @Failure 400 {object} dtos.APIError "Invalid credentials or request body"
// @Failure 500 {object} dtos.APIError "Internal server error"
func (c *UserController) LoginUser(ginCtx *gin.Context) {
	var req dtos.LoginUserRequest
	if err := utils.ValidateRequestBody(ginCtx, &req); err != nil {
		ginCtx.Error(err)
		return
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
