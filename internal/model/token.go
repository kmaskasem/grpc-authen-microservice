package model

import "time"

type TokenBlacklist struct {
	Token    string    `bson:"token"`
	ExpireAt time.Time `bson:"expired_at"`
}
