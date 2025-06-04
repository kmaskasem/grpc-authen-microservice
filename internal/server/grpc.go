package server

import (
	"context"
	"log"
	"net"

	"github.com/kmaskasem/grpc-authen-microservice/config"
	"github.com/kmaskasem/grpc-authen-microservice/internal/database"
	grpchandler "github.com/kmaskasem/grpc-authen-microservice/internal/handler/grpc"
	"github.com/kmaskasem/grpc-authen-microservice/internal/repository"
	"github.com/kmaskasem/grpc-authen-microservice/internal/service"
	"github.com/kmaskasem/grpc-authen-microservice/utils"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	pb "github.com/kmaskasem/grpc-authen-microservice/proto"
	"google.golang.org/grpc"
)

func StartGRPCServer() {
	cfg := config.LoadConfig()

	client, err := database.ConnectMongoDB(cfg.MongoURI)
	if err != nil {
		log.Fatal("Failed to connect to MongoDB:", err)
	}

	// Init secret
	utils.Init(cfg.JWTSecret)

	db := client.Database(cfg.MongoDB)
	userRepo := repository.NewUserRepository(db)
	tokenRepo := repository.NewTokenRepository(db)
	loginAttemptRepo := repository.NewLoginAttemptRepository(db)

	loginAttemptRepo.Collection.Indexes().CreateOne(context.Background(), mongo.IndexModel{
		Keys:    bson.M{"timestamp": 1},
		Options: options.Index().SetExpireAfterSeconds(60),
	})

	authService := service.NewAuthService(tokenRepo, userRepo, loginAttemptRepo)
	userService := service.NewUserService(userRepo)

	// Start listener
	lis, err := net.Listen("tcp", cfg.GRPCPort)
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	// gRPC Server
	grpcServer := grpc.NewServer(
		grpc.UnaryInterceptor(utils.AuthInterceptor(cfg.JWTSecret, userRepo, tokenRepo)),
	)

	// Register Services
	pb.RegisterAuthServiceServer(grpcServer, &grpchandler.AuthHandler{AuthService: authService})
	pb.RegisterUserServiceServer(grpcServer, &grpchandler.UserHandler{UserService: userService})

	log.Println("gRPC server started on", cfg.GRPCPort)
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
