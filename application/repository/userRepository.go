package repository

import (
	"context"
	"errors"

	"github.com/joaovds/chat/application/domain"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserRepository struct {
	Collection *mongo.Collection
}

func NewUserRepository(db *mongo.Database) *UserRepository {
	return &UserRepository{
		Collection: db.Collection("users"),
	}
}

func (ur *UserRepository) GetUserByNickName(nickName string) (*domain.User, error) {
	filter := bson.M{"nickname": nickName}
	var user domain.User
	result := ur.Collection.FindOne(context.TODO(), filter)
	if err := result.Err(); err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, errors.New("User not found")
		}
		return nil, err
	}
	if err := result.Decode(&user); err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, errors.New("User not found")
		}
		return nil, err
	}
	return &user, nil
}

func (ur *UserRepository) CreateUser(user *domain.User) error {
	_, err := ur.Collection.InsertOne(context.TODO(), user)
	return err
}
