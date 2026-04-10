package services

import (
	"context"
	"fmt"
	"go-api/src/entities"
	"go-api/src/utils"
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

type UserService struct {
	repository userRepository
}

func NewUserService(repository userRepository) *UserService {
	return &UserService{repository: repository}
}

func (s *UserService) RegisterUser(ctx context.Context, email, password string) error {
	existing, err := s.repository.FindByEmail(ctx, email)
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

	return s.repository.Create(ctx, user)
}
