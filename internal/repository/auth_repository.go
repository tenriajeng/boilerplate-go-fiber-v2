package repository

import (
	"context"
	"errors"
	"time"

	"boilerplate-go-fiber-v2/internal/domain/entity"
	"boilerplate-go-fiber-v2/internal/domain/repository"

	"gorm.io/gorm"
)

type authRepository struct {
	db *gorm.DB
}

// NewAuthRepository creates a new auth repository
func NewAuthRepository(db *gorm.DB) repository.AuthRepository {
	return &authRepository{db: db}
}

// Session management methods

// CreateSession creates a new auth session
func (r *authRepository) CreateSession(ctx context.Context, session *entity.AuthSession) error {
	return r.db.WithContext(ctx).Create(session).Error
}

// GetSessionByToken gets a session by token
func (r *authRepository) GetSessionByToken(ctx context.Context, token string) (*entity.AuthSession, error) {
	var session entity.AuthSession
	err := r.db.WithContext(ctx).Where("token = ?", token).First(&session).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("session not found")
		}
		return nil, err
	}
	return &session, nil
}

// GetSessionByRefreshToken gets a session by refresh token
func (r *authRepository) GetSessionByRefreshToken(ctx context.Context, refreshToken string) (*entity.AuthSession, error) {
	var session entity.AuthSession
	err := r.db.WithContext(ctx).Where("refresh_token = ?", refreshToken).First(&session).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("session not found")
		}
		return nil, err
	}
	return &session, nil
}

// DeleteSession deletes a session by token
func (r *authRepository) DeleteSession(ctx context.Context, token string) error {
	return r.db.WithContext(ctx).Where("token = ?", token).Delete(&entity.AuthSession{}).Error
}

// DeleteSessionsByUserID deletes all sessions for a user
func (r *authRepository) DeleteSessionsByUserID(ctx context.Context, userID uint) error {
	return r.db.WithContext(ctx).Where("user_id = ?", userID).Delete(&entity.AuthSession{}).Error
}

// CleanExpiredSessions removes expired sessions
func (r *authRepository) CleanExpiredSessions(ctx context.Context) error {
	return r.db.WithContext(ctx).Where("expires_at < ?", time.Now()).Delete(&entity.AuthSession{}).Error
}

// Password reset methods

// CreatePasswordReset creates a new password reset
func (r *authRepository) CreatePasswordReset(ctx context.Context, reset *entity.PasswordReset) error {
	return r.db.WithContext(ctx).Create(reset).Error
}

// GetPasswordResetByToken gets a password reset by token
func (r *authRepository) GetPasswordResetByToken(ctx context.Context, token string) (*entity.PasswordReset, error) {
	var reset entity.PasswordReset
	err := r.db.WithContext(ctx).Where("token = ?", token).First(&reset).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("password reset not found")
		}
		return nil, err
	}
	return &reset, nil
}

// MarkPasswordResetUsed marks a password reset as used
func (r *authRepository) MarkPasswordResetUsed(ctx context.Context, token string) error {
	return r.db.WithContext(ctx).Model(&entity.PasswordReset{}).Where("token = ?", token).Update("used", true).Error
}

// CleanExpiredPasswordResets removes expired password resets
func (r *authRepository) CleanExpiredPasswordResets(ctx context.Context) error {
	return r.db.WithContext(ctx).Where("expires_at < ?", time.Now()).Delete(&entity.PasswordReset{}).Error
}

// TFA code methods

// CreateTFACode creates a new TFA code
func (r *authRepository) CreateTFACode(ctx context.Context, code *entity.TFACode) error {
	return r.db.WithContext(ctx).Create(code).Error
}

// GetTFACodeByCode gets a TFA code by code
func (r *authRepository) GetTFACodeByCode(ctx context.Context, code string) (*entity.TFACode, error) {
	var tfaCode entity.TFACode
	err := r.db.WithContext(ctx).Where("code = ?", code).First(&tfaCode).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("TFA code not found")
		}
		return nil, err
	}
	return &tfaCode, nil
}

// MarkTFACodeUsed marks a TFA code as used
func (r *authRepository) MarkTFACodeUsed(ctx context.Context, code string) error {
	return r.db.WithContext(ctx).Model(&entity.TFACode{}).Where("code = ?", code).Update("used", true).Error
}

// UpdateSession updates a session
func (r *authRepository) UpdateSession(ctx context.Context, session *entity.AuthSession) error {
	return r.db.WithContext(ctx).Save(session).Error
}

// CleanExpiredTFACodes removes expired TFA codes
func (r *authRepository) CleanExpiredTFACodes(ctx context.Context) error {
	return r.db.WithContext(ctx).Where("expires_at < ?", time.Now()).Delete(&entity.TFACode{}).Error
}
