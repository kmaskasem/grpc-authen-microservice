package main

import (
	"log"
	"net"

	"github.com/kmaskasem/grpc-authen-microservice/config"
	"github.com/kmaskasem/grpc-authen-microservice/internal/database"
	"github.com/kmaskasem/grpc-authen-microservice/internal/repository"
	"github.com/kmaskasem/grpc-authen-microservice/internal/service"
	"github.com/kmaskasem/grpc-authen-microservice/utils"

	grpchandler "github.com/kmaskasem/grpc-authen-microservice/internal/handler/grpc"
	pb "github.com/kmaskasem/grpc-authen-microservice/proto"
	"google.golang.org/grpc"
)

func main() {

	cfg := config.LoadConfig()

	client, err := database.ConnectMongoDB(cfg.MongoURI)
	if err != nil {
		log.Fatal("Can not Connect to MongoDB:", err)
	}

	//Init JWT_SECRET
	utils.Init(cfg.JWTSecret)

	db := client.Database(cfg.MongoDB)
	userRepo := repository.NewUserRepository(db)
	tokenRepo := repository.NewTokenRepository(db)
	authService := service.NewAuthService(tokenRepo, userRepo)
	// userService := service.NewUserService(userRepo)
	// log.Println("Loaded config:", cfg)

	// gRPC server
	lis, err := net.Listen("tcp", ":50051") // หรือพอร์ตอื่น
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterAuthServiceServer(grpcServer, &grpchandler.AuthHandler{AuthService: authService})

	log.Println("gRPC server started on :50051")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
	// user := &model.User{
	// 	Name:     "James3",
	// 	Email:    "J@example3.com",
	// 	Password: "Jpassword",
	// }

	// err = authService.Register(ctx, user)
	// if err != nil {
	// 	log.Fatal("Register failed:", err)
	// }

	// log.Println("User registered successfully")
}
