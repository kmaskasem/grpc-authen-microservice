package service

import (
	"github.com/kmaskasem/grpc-authen-microservice/internal/repository"
)

type UserService struct {
	// TokenRepo *repository.TokenRepository
	UserRepo *repository.UserRepository
}

func NewUserService(tokenRepo *repository.TokenRepository, userRepo *repository.UserRepository) *UserService {
	return &UserService{
		// TokenRepo: tokenRepo,
		UserRepo: userRepo,
	}
}

// func (s *AuthService) ListUsers(ctx context.Context, token string) error {

// }

// func (s *UserService) GetProfile(ctx context.Context, user *model.User) error {
// }

// func (s *AuthService) UpdateProfile(ctx context.Context, email string, password string) (string, error) {
// }

// func (s *AuthService) DeleteProfile(ctx context.Context, token string) error {
// }
