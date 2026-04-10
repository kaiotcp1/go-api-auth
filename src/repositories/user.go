package repositories

import (
	"context"
	"go-api/src/entities"
	"go-api/src/utils"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserRepository struct {
	collection *mongo.Collection
}

func NewUserRepository(collection *mongo.Collection) *UserRepository {
	return &UserRepository{collection: collection}
}

func (r *UserRepository) FindByEmail(ctx context.Context, email string) (*entities.User, error) {
	var user entities.User
	err := r.collection.FindOne(ctx, bson.M{"email": email}).Decode(&user)
	if err == mongo.ErrNoDocuments {
		return nil, nil
	}
	if err != nil {
		return nil, utils.InternalServerError("erro ao buscar usuário")
	}
	return &user, nil
}

func (r *UserRepository) Create(ctx context.Context, user *entities.User) error {
	_, err := r.collection.InsertOne(ctx, user)
	if err != nil {
		return utils.InternalServerError("erro ao criar usuário")
	}
	return nil
}
