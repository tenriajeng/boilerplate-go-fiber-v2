package model

import (
	"time"

	"boilerplate-go-fiber-v2/internal/domain/entity"
)

type AuthSessionModel struct {
	ID           uint      `gorm:"primaryKey;autoIncrement"`
	UserID       uint      `gorm:"not null"`
	Token        string    `gorm:"uniqueIndex;not null"`
	RefreshToken string    `gorm:"uniqueIndex;not null"`
	ExpiresAt    time.Time `gorm:"not null"`
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

type PasswordResetModel struct {
	ID        uint      `gorm:"primaryKey;autoIncrement"`
	UserID    uint      `gorm:"not null"`
	Token     string    `gorm:"uniqueIndex;not null"`
	ExpiresAt time.Time `gorm:"not null"`
	Used      bool      `gorm:"default:false"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

type TFACodeModel struct {
	ID        uint      `gorm:"primaryKey;autoIncrement"`
	UserID    uint      `gorm:"not null"`
	Code      string    `gorm:"not null"`
	ExpiresAt time.Time `gorm:"not null"`
	Used      bool      `gorm:"default:false"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

type EmailVerificationModel struct {
	ID         uint      `gorm:"primaryKey;autoIncrement"`
	UserID     uint      `gorm:"not null"`
	Email      string    `gorm:"not null"`
	Token      string    `gorm:"uniqueIndex;not null"`
	ExpiresAt  time.Time `gorm:"not null"`
	VerifiedAt *time.Time
	CreatedAt  time.Time
	UpdatedAt  time.Time
}

func (AuthSessionModel) TableName() string {
	return "auth_sessions"
}

func (PasswordResetModel) TableName() string {
	return "password_resets"
}

func (TFACodeModel) TableName() string {
	return "tfa_codes"
}

func (EmailVerificationModel) TableName() string {
	return "email_verifications"
}

// AuthSession conversion methods
func (m *AuthSessionModel) ToEntity() *entity.AuthSession {
	return &entity.AuthSession{
		ID:           m.ID,
		UserID:       m.UserID,
		Token:        m.Token,
		RefreshToken: m.RefreshToken,
		ExpiresAt:    m.ExpiresAt,
		CreatedAt:    m.CreatedAt,
		UpdatedAt:    m.UpdatedAt,
	}
}

func (m *AuthSessionModel) FromEntity(session *entity.AuthSession) {
	m.ID = session.ID
	m.UserID = session.UserID
	m.Token = session.Token
	m.RefreshToken = session.RefreshToken
	m.ExpiresAt = session.ExpiresAt
	m.CreatedAt = session.CreatedAt
	m.UpdatedAt = session.UpdatedAt
}

// PasswordReset conversion methods
func (m *PasswordResetModel) ToEntity() *entity.PasswordReset {
	return &entity.PasswordReset{
		ID:        m.ID,
		UserID:    m.UserID,
		Token:     m.Token,
		ExpiresAt: m.ExpiresAt,
		Used:      m.Used,
		CreatedAt: m.CreatedAt,
		UpdatedAt: m.UpdatedAt,
	}
}

func (m *PasswordResetModel) FromEntity(reset *entity.PasswordReset) {
	m.ID = reset.ID
	m.UserID = reset.UserID
	m.Token = reset.Token
	m.ExpiresAt = reset.ExpiresAt
	m.Used = reset.Used
	m.CreatedAt = reset.CreatedAt
	m.UpdatedAt = reset.UpdatedAt
}

// TFACode conversion methods
func (m *TFACodeModel) ToEntity() *entity.TFACode {
	return &entity.TFACode{
		ID:        m.ID,
		UserID:    m.UserID,
		Code:      m.Code,
		ExpiresAt: m.ExpiresAt,
		Used:      m.Used,
		CreatedAt: m.CreatedAt,
		UpdatedAt: m.UpdatedAt,
	}
}

func (m *TFACodeModel) FromEntity(tfa *entity.TFACode) {
	m.ID = tfa.ID
	m.UserID = tfa.UserID
	m.Code = tfa.Code
	m.ExpiresAt = tfa.ExpiresAt
	m.Used = tfa.Used
	m.CreatedAt = tfa.CreatedAt
	m.UpdatedAt = tfa.UpdatedAt
}

// EmailVerification conversion methods
func (m *EmailVerificationModel) ToEntity() *entity.EmailVerification {
	return &entity.EmailVerification{
		ID:         m.ID,
		UserID:     m.UserID,
		Email:      m.Email,
		Token:      m.Token,
		ExpiresAt:  m.ExpiresAt,
		VerifiedAt: m.VerifiedAt,
		CreatedAt:  m.CreatedAt,
		UpdatedAt:  m.UpdatedAt,
	}
}

func (m *EmailVerificationModel) FromEntity(verification *entity.EmailVerification) {
	m.ID = verification.ID
	m.UserID = verification.UserID
	m.Email = verification.Email
	m.Token = verification.Token
	m.ExpiresAt = verification.ExpiresAt
	m.VerifiedAt = verification.VerifiedAt
	m.CreatedAt = verification.CreatedAt
	m.UpdatedAt = verification.UpdatedAt
}
