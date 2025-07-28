package repository

import (
	"boilerplate-go-fiber-v2/internal/domain/entity"
	"context"
)

type AuthRepository interface {
	// Session management
	CreateSession(ctx context.Context, session *entity.AuthSession) error
	GetSessionByToken(ctx context.Context, token string) (*entity.AuthSession, error)
	GetSessionByRefreshToken(ctx context.Context, refreshToken string) (*entity.AuthSession, error)
	UpdateSession(ctx context.Context, session *entity.AuthSession) error
	DeleteSession(ctx context.Context, token string) error
	DeleteSessionsByUserID(ctx context.Context, userID uint) error
	CleanExpiredSessions(ctx context.Context) error

	// Password reset
	CreatePasswordReset(ctx context.Context, reset *entity.PasswordReset) error
	GetPasswordResetByToken(ctx context.Context, token string) (*entity.PasswordReset, error)
	MarkPasswordResetUsed(ctx context.Context, token string) error
	CleanExpiredPasswordResets(ctx context.Context) error

	// TFA codes
	CreateTFACode(ctx context.Context, code *entity.TFACode) error
	GetTFACodeByCode(ctx context.Context, code string) (*entity.TFACode, error)
	MarkTFACodeUsed(ctx context.Context, code string) error
	CleanExpiredTFACodes(ctx context.Context) error
}
