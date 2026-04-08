package repositories

import (
	"context"
	"fmt"
	"go-api/src/utils"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoRepositoryContext struct {
	*mongo.Collection //embedded field
	client            *mongo.Client
}

func NewMongoRepositoryContext(uri, dbName, collectionName string) (*MongoRepositoryContext, error) {
	clientOptions := options.Client().ApplyURI(uri)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		return nil, utils.InternalServerError(fmt.Sprintf("Erro ao conectar ao banco de dados: %v", err))
	}

	err = client.Ping(context.Background(), nil)
	if err != nil {
		return nil, utils.InternalServerError(fmt.Sprintf("Erro ao conectar ao banco de dados: %v", err))
	}

	collection := client.Database(dbName).Collection(collectionName)

	return &MongoRepositoryContext{
		Collection: collection,
		client:     client,
	}, nil

}

func (r *MongoRepositoryContext) Create(contextServer context.Context, document interface{}) error {
	_, err := r.Collection.InsertOne(contextServer, document)
	if err != nil {
		return utils.BadRequestError(fmt.Sprintf("Erro ao inserir documento: %v", err))
	}
	return nil
}
