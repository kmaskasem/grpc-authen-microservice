package grpchandler

import (
	"context"

	"github.com/kmaskasem/grpc-authen-microservice/internal/model"
	"github.com/kmaskasem/grpc-authen-microservice/internal/service"

	pb "github.com/kmaskasem/grpc-authen-microservice/proto"
)

type AuthHandler struct {
	pb.UnimplementedAuthServiceServer
	AuthService *service.AuthService
}

// Register Handler
func (s *AuthHandler) Register(ctx context.Context, req *pb.RegisterRequest) (*pb.RegisterResponse, error) {
	err := s.AuthService.Register(ctx, &model.User{
		Email:    req.Email,
		Password: req.Password,
		Name:     req.Name,
	})
	if err != nil {
		return nil, err
	}
	return &pb.RegisterResponse{Message: "registered"}, nil
}

// Login Handler
func (s *AuthHandler) Login(ctx context.Context, req *pb.LoginRequest) (*pb.LoginResponse, error) {
	token, err := s.AuthService.Login(ctx, req.Email, req.Password)
	if err != nil {
		return nil, err
	}

	return &pb.LoginResponse{Token: token}, nil
}

// Logout Handler
func (s *AuthHandler) Logout(ctx context.Context, req *pb.LogoutRequest) (*pb.LogoutResponse, error) {
	err := s.AuthService.Logout(ctx)
	if err != nil {
		return nil, err
	}

	return &pb.LogoutResponse{Message: "log out"}, nil
}
