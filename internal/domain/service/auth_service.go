package service

import (
	"context"

	"boilerplate-go-fiber-v2/internal/domain/entity"
	"boilerplate-go-fiber-v2/pkg/jwt"
)

type AuthService interface {
	Login(ctx context.Context, email, password string) (*entity.User, *entity.AuthSession, error)
	Logout(ctx context.Context, token string) error
	RefreshToken(ctx context.Context, refreshToken string) (*entity.AuthSession, error)
	Register(ctx context.Context, user *entity.User) error
	ValidateToken(ctx context.Context, token string) (*jwt.Claims, error)
	CreatePasswordReset(ctx context.Context, email string) error
	ResetPassword(ctx context.Context, token, newPassword string) error
	CreateTFACode(ctx context.Context, userID uint) error
	VerifyTFACode(ctx context.Context, userID uint, code string) error
	EnableTFA(ctx context.Context, userID uint) (*entity.User, error)
	DisableTFA(ctx context.Context, userID uint) error
	VerifyTFA(ctx context.Context, userID uint, code string) error
}
