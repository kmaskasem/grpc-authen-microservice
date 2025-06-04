package grpchandler

import (
	"context"

	"github.com/kmaskasem/grpc-authen-microservice/internal/service"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	pb "github.com/kmaskasem/grpc-authen-microservice/proto"
)

type UserHandler struct {
	pb.UnimplementedUserServiceServer
	UserService *service.UserService
}

func (h *UserHandler) ListUsers(ctx context.Context, req *pb.ListUsersRequest) (*pb.ListUsersResponse, error) {

	users, err := h.UserService.ListUsers(ctx, req.Name, req.Email, int64(req.Page), int64(req.Limit))
	if err != nil {
		return nil, err
	}
	var result []*pb.User
	for _, u := range users {
		result = append(result, &pb.User{
			Id:    u.ID.Hex(),
			Name:  u.Name,
			Email: u.Email,
		})
	}
	return &pb.ListUsersResponse{Users: result}, nil
}

func (h *UserHandler) GetProfile(ctx context.Context, req *pb.GetProfileRequest) (*pb.GetProfileResponse, error) {
	userId, ok := ctx.Value("userId").(string)
	if !ok || userId == "" {
		return nil, status.Error(codes.Unauthenticated, "unauthenticated")
	}

	user, err := h.UserService.GetProfile(ctx, userId)
	if err != nil {
		return nil, err
	}
	return &pb.GetProfileResponse{
		User: &pb.User{
			Id:    user.ID.Hex(),
			Name:  user.Name,
			Email: user.Email,
		},
	}, nil
}

func (h *UserHandler) UpdateProfile(ctx context.Context, req *pb.UpdateProfileRequest) (*pb.UpdateProfileResponse, error) {
	userId, ok := ctx.Value("userId").(string)
	if !ok || userId == "" {
		return nil, status.Error(codes.Unauthenticated, "unauthenticated")
	}

	err := h.UserService.UpdateProfile(ctx, userId, req.Name, req.Email)
	if err != nil {
		return nil, err
	}
	return &pb.UpdateProfileResponse{Message: "updated"}, nil
}

func (h *UserHandler) DeleteProfile(ctx context.Context, req *pb.DeleteProfileRequest) (*pb.DeleteProfileResponse, error) {
	userId, ok := ctx.Value("userId").(string)
	if !ok || userId == "" {
		return nil, status.Error(codes.Unauthenticated, "unauthenticated")
	}

	err := h.UserService.DeleteProfile(ctx, userId)
	if err != nil {
		return nil, err
	}
	return &pb.DeleteProfileResponse{Message: "deleted"}, nil
}
