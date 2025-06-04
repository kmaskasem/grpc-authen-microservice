package model

import "time"

type LoginAttempt struct {
	Email     string    `bson:"email"`
	Timestamp time.Time `bson:"timestamp"`
}
