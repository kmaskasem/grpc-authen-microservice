package repository

import (
	"context"
	"log"

	"github.com/kmaskasem/grpc-authen-microservice/internal/model"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
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

func (r *UserRepository) ListUsers(ctx context.Context, filter bson.M, page, limit int64) ([]*model.User, error) {
	opts := options.Find().SetSkip((page - 1) * limit).SetLimit(limit)
	cursor, err := r.Collection.Find(ctx, filter, opts)
	if err != nil {
		return nil, err
	}
	var users []*model.User
	if err = cursor.All(ctx, &users); err != nil {
		return nil, err
	}
	return users, nil
}

func (r *UserRepository) FindByID(ctx context.Context, id primitive.ObjectID) (*model.User, error) {
	var user model.User

	err := r.Collection.FindOne(ctx, bson.M{"_id": id, "deleted": false}).Decode(&user)
	if err != nil {
		return nil, err
	}
	return &user, err
}

func (r *UserRepository) UpdateUser(ctx context.Context, id primitive.ObjectID, update bson.M) error {
	_, err := r.Collection.UpdateOne(ctx, bson.M{"_id": id}, bson.M{"$set": update})
	return err
}

func (r *UserRepository) SoftDeleteUser(ctx context.Context, id primitive.ObjectID) error {
	_, err := r.Collection.UpdateOne(ctx, bson.M{"_id": id}, bson.M{"$set": bson.M{"deleted": true}})
	return err
}
