package service

import (
	"context"
	"errors"
	"time"

	"github.com/kmaskasem/grpc-authen-microservice/internal/model"
	"github.com/kmaskasem/grpc-authen-microservice/internal/repository"
	"github.com/kmaskasem/grpc-authen-microservice/utils"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

type AuthService struct {
	TokenRepo *repository.TokenRepository
	UserRepo  *repository.UserRepository
}

func NewAuthService(tokenRepo *repository.TokenRepository, userRepo *repository.UserRepository) *AuthService {
	return &AuthService{
		TokenRepo: tokenRepo,
		UserRepo:  userRepo,
	}
}

func (s *AuthService) Register(ctx context.Context, user *model.User) error {
	existingUser, _ := s.UserRepo.FindByEmail(ctx, user.Email)
	if existingUser != nil {
		return errors.New("email already exists")
	}

	if len(user.Password) < 6 {
		return errors.New("password must be at least 6 characters")
	}

	hashPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	user.Password = string(hashPassword)

	err = s.UserRepo.CreateUser(ctx, user)
	if err != nil {
		return err
	}

	return nil
}

func (s *AuthService) Login(ctx context.Context, email string, password string) (string, error) {
	user, err := s.UserRepo.FindByEmail(ctx, email)
	if err != nil {
		return "", errors.New("email is not correct")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return "", errors.New("password is not correct")
	}

	token, err := utils.GenerateJWT(user.ID.Hex())
	if err != nil {
		return "", err
	}

	return token, nil
}

func (s *AuthService) Logout(ctx context.Context, token string) error {
	//Check token with signature
	parsedToken, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		return utils.GetJWTSecret(), nil
	})
	if err != nil {
		return err
	}

	claims, ok := parsedToken.Claims.(jwt.MapClaims)
	if !ok || !parsedToken.Valid {
		return errors.New("invalid token")
	}

	expUnix := int64(claims["exp"].(float64))
	expTime := time.Unix(expUnix, 0)

	return s.TokenRepo.BlacklistToken(ctx, token, expTime)
}
