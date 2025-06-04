package repository

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type LoginAttemptRepository struct {
	Collection *mongo.Collection
}

func NewLoginAttemptRepository(db *mongo.Database) *LoginAttemptRepository {
	return &LoginAttemptRepository{
		Collection: db.Collection("login_attempts"),
	}
}

func (r *LoginAttemptRepository) CountRecentAttempts(ctx context.Context, email string) (int64, error) {
	filter := bson.M{"email": email}
	return r.Collection.CountDocuments(ctx, filter)
}

func (r *LoginAttemptRepository) AddAttempt(ctx context.Context, email string) error {
	_, err := r.Collection.InsertOne(ctx, bson.M{
		"email":     email,
		"timestamp": time.Now(),
	})
	return err
}
