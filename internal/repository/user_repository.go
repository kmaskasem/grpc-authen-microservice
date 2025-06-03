package repository

import (
	"context"
	"log"

	"github.com/kmaskasem/grpc-authen-microservice/internal/model"

	"go.mongodb.org/mongo-driver/mongo"
)

type UserRepository struct {
	Collection *mongo.Collection
}

// Constructor UserRepository
func NewUserRepository(db *mongo.Database) *UserRepository {
	return &UserRepository{
		Collection: db.Collection("user"),
	}
}

func (r *UserRepository) CreateUser(ctx context.Context, user *model.User) error {
	result, err := r.Collection.InsertOne(ctx, user)
	if err != nil {
		return err
	}
	log.Println("User Created with ID:", result.InsertedID)
	return err
}

func (r *UserRepository) FindByEmail(ctx context.Context, email string) (*model.User, error) {
	var user model.User
	err := r.Collection.FindOne(ctx, map[string]interface{}{"email": email}).Decode(&user)
	if err != nil {
		return nil, err
	}
	return &user, nil
}
