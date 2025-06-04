package service

import (
	"context"

	"github.com/kmaskasem/grpc-authen-microservice/internal/model"
	"github.com/kmaskasem/grpc-authen-microservice/internal/repository"
	"github.com/kmaskasem/grpc-authen-microservice/utils"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type UserService struct {
	Repo *repository.UserRepository
}

func NewUserService(userRepo *repository.UserRepository) *UserService {
	return &UserService{
		Repo: userRepo,
	}
}

func (s *UserService) ListUsers(ctx context.Context, name, email string, page, limit int64) ([]*model.User, error) {
	filter := bson.M{"deleted": false}
	if name != "" {
		filter["name"] = bson.M{"$regex": name, "$options": "i"}
	}
	if email != "" {
		filter["email"] = bson.M{"$regex": email, "$options": "i"}
	}
	return s.Repo.ListUsers(ctx, filter, page, limit)
}

func (s *UserService) GetProfile(ctx context.Context, id string) (*model.User, error) {
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}
	return s.Repo.FindByID(ctx, objID)
}

func (s *UserService) UpdateProfile(ctx context.Context, id, name, email string) error {
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	// Validate Email format
	if !utils.ValidateEmail(email) {
		return status.Error(codes.InvalidArgument, "invalid email format")
	}

	// Check if email is used by another user
	existingUser, _ := s.Repo.FindByEmail(ctx, email)
	if existingUser != nil && existingUser.ID != objID {
		return status.Error(codes.AlreadyExists, "email already in use by another user")
	}

	return s.Repo.UpdateUser(ctx, objID, bson.M{"name": name, "email": email})
}

func (s *UserService) DeleteProfile(ctx context.Context, id string) error {
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}
	return s.Repo.SoftDeleteUser(ctx, objID)
}
