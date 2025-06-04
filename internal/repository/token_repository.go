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

// Constructor TokenRepository
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

func (r *TokenRepository) IsBlacklisted(ctx context.Context, token string) (bool, error) {
	count, err := r.Collection.CountDocuments(ctx, bson.M{"token": token})
	return count > 0, err
}
