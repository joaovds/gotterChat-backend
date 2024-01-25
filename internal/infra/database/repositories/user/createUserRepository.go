package user_repository

import (
	"context"

	"github.com/joaovds/chat/internal/data/dtos"
	"go.mongodb.org/mongo-driver/mongo"
)

type CreateUserRepository struct {
	Collection *mongo.Collection
}

func NewCreateUserRepository(db *mongo.Database) *CreateUserRepository {
	return &CreateUserRepository{
		Collection: db.Collection("users"),
	}
}

func (c *CreateUserRepository) Create(user *dtos.CreateUserInputDTO) error {
	_, err := c.Collection.InsertOne(context.TODO(), user)

	return err
}
