package services

import (
	"context"
	"fmt"
	"go-api/src/entities"
	"go-api/src/utils"
	"go-api/src/utils/middleware"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
)

// userRepository é a interface que o serviço precisa — definida aqui,
// satisfeita implicitamente por *repositories.UserRepository.
type userRepository interface {
	FindByEmail(ctx context.Context, email string) (*entities.User, error)
	Create(ctx context.Context, user *entities.User) error
}

// Struct UserService implementa a interface UserRepository
type UserService struct {
	repository userRepository
}

func NewUserService(repository userRepository) *UserService {
	return &UserService{repository: repository}
}

func (service *UserService) RegisterUser(ctx context.Context, email, password string) error {
	existing, err := service.repository.FindByEmail(ctx, email)
	if err != nil {
		return err
	}
	if existing != nil {
		return utils.ConflictError(fmt.Sprintf("%s já cadastrado.", email))
	}

	hashed, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return utils.InternalServerError("erro ao processar senha")
	}

	user := &entities.User{
		ID:        primitive.NewObjectID(),
		Email:     email,
		Password:  string(hashed),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	return service.repository.Create(ctx, user)
}

func (service *UserService) LoginUser(ctx context.Context, email, password string) (string, error) {

	user, err := service.repository.FindByEmail(ctx, email)
	if err != nil {
		return "", err
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))

	if err != nil {
		return "", utils.BadRequestError("email ou senha incorretos")
	}

	return middleware.GenerateToken(user.Email, user.ID)

}
