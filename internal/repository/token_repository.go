package repository

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type TokenRepository struct {
	Collection *mongo.Collection
}

// Constructor UserRepository
func NewTokenRepository(db *mongo.Database) *TokenRepository {
	return &TokenRepository{
		Collection: db.Collection("token"),
	}
}

func (r *TokenRepository) BlacklistToken(ctx context.Context, token string, expireAt time.Time) error {
	_, err := r.Collection.InsertOne(ctx, bson.M{
		"token":      token,
		"expired_at": expireAt,
	})
	if err != nil {
		return err
	}

	return err
}
