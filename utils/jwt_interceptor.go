package utils

import (
	"context"
	"strings"

	"github.com/golang-jwt/jwt/v5"
	"github.com/kmaskasem/grpc-authen-microservice/internal/repository"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

func AuthInterceptor(secretKey string, repo *repository.UserRepository, tokenRepo *repository.TokenRepository) grpc.UnaryServerInterceptor {
	return func(
		ctx context.Context,
		req interface{},
		info *grpc.UnaryServerInfo,
		handler grpc.UnaryHandler,
	) (interface{}, error) {
		// Skip Login and Register API
		if strings.Contains(info.FullMethod, "Login") || strings.Contains(info.FullMethod, "Register") {
			return handler(ctx, req)
		}

		// Get metadata from context (Header)
		md, ok := metadata.FromIncomingContext(ctx)
		if !ok {
			return nil, status.Error(codes.Unauthenticated, "missing metadata")
		}

		// Check md["authorization"] (Token) has a value
		tokenStrs := md["authorization"]
		if len(tokenStrs) == 0 {
			return nil, status.Error(codes.Unauthenticated, "missing token")
		}

		// Get Token
		tokenStr := strings.TrimPrefix(tokenStrs[0], "Bearer ")

		// Convert Token to Object
		token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
			return []byte(secretKey), nil
		})

		// Check token with signature
		if err != nil || !token.Valid {
			return nil, status.Error(codes.Unauthenticated, "invalid token")
		}

		// Check is Token already logout
		if tokenRepo != nil {
			isBlacklisted, err := tokenRepo.IsBlacklisted(ctx, tokenStr)
			if err != nil {
				return nil, status.Error(codes.Internal, "error checking token")
			}
			if isBlacklisted {
				return nil, status.Error(codes.Unauthenticated, "token has been revoked")
			}
		}

		// token is pass check
		if claims, ok := token.Claims.(jwt.MapClaims); ok {
			// get user_id
			if userIDStr, ok := claims["user_id"].(string); ok {
				// Convert user_id to ObjectID
				userID, err := primitive.ObjectIDFromHex(userIDStr)
				if err != nil {
					return nil, status.Error(codes.InvalidArgument, "invalid user id")
				}

				// Check is user deleted?
				user, err := repo.FindByID(ctx, userID)
				if err != nil || user.Deleted {
					return nil, status.Error(codes.PermissionDenied, "user deleted or not found")
				}

				// Pack userId to context to use on handler next
				ctx = context.WithValue(ctx, "userId", userIDStr)
			}
		}

		// return handler if all pass
		return handler(ctx, req)
	}
}
