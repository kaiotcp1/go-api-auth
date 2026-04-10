package services

import (
	"fmt"
	"go-api/src/dtos"
	"go-api/src/repositories"
	"go-api/src/utils"
	"time"

	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	repository *repositories.UserRepository
}

func NewUserService(repository *repositories.UserRepository) *UserService {
	return &UserService{repository: repository}
}

func (service *UserService) CreateUser(email, password string) error {

	hashedPassword, err := HashPassword(password)
	if err != nil {
		return err
	}

	password = hashedPassword

	user := dtos.User{
		Email:     email,
		Password:  hashedPassword,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	contextServer := utils.CreateContextServerWithTimeout()

	err = service.repository.Create(contextServer, user)
	if err != nil {
		return err
	}

	return nil
}

func HashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", utils.BadRequestError(fmt.Sprintf("Erro ao gerar hash: %v", err))
	}
	return string(hashedPassword), nil
}
