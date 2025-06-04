package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	MongoURI  string
	MongoDB   string
	JWTSecret string
	GRPCPort  string
}

func LoadConfig() *Config {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file.")
	}

	return &Config{
		MongoURI:  os.Getenv("MONGO_URI"),
		MongoDB:   os.Getenv("MONGO_DB"),
		JWTSecret: os.Getenv("JWT_SECRET"),
		GRPCPort:  os.Getenv("GRPC_PORT"),
	}
}
