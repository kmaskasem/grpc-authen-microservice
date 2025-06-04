package service

import (
	"context"
	"errors"
	"time"

	"github.com/kmaskasem/grpc-authen-microservice/internal/model"
	"github.com/kmaskasem/grpc-authen-microservice/internal/repository"
	"github.com/kmaskasem/grpc-authen-microservice/utils"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

type AuthService struct {
	TokenRepo *repository.TokenRepository
	UserRepo  *repository.UserRepository
}

// Constructer AuthService
func NewAuthService(tokenRepo *repository.TokenRepository, userRepo *repository.UserRepository) *AuthService {
	return &AuthService{
		TokenRepo: tokenRepo,
		UserRepo:  userRepo,
	}
}

func (s *AuthService) Register(ctx context.Context, user *model.User) error {
	// Validate Email format
	if !utils.ValidateEmail(user.Email) {
		return status.Error(codes.InvalidArgument, "invalid email format")
	}

	// Validate is Email used
	existingUser, _ := s.UserRepo.FindByEmail(ctx, user.Email)
	if existingUser != nil {
		return errors.New("email already exists")
	}

	// Validate Password strength
	if err := utils.ValidatePassword(user.Password); err != nil {
		return err
	}

	// Hash Password
	hashPassword, err := utils.HashPassword(user.Password)
	if err != nil {
		return err
	}
	user.Password = hashPassword
	user.Deleted = false

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

func (s *AuthService) Logout(ctx context.Context) error {
	token := ctx.Value("token").(string)

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
