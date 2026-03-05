package services

import (
	"go-api/src/dtos"
	"go-api/src/repositories"
	"go-api/src/utils"
)

type UserService struct {
	repository *repositories.UserRepository
}

func NewUserService(repository *repositories.UserRepository) *UserService {
	return &UserService{repository: repository}
}

func (service *UserService) CreateUser(email, password string) error {
	user := dtos.User{
		Email:    email,
		Password: password,
	}

	contextServer := utils.CreateContextServerWithTimeout()

	err := service.repository.Create(contextServer, user)
	if err != nil {
		return err
	}

	return nil
}
