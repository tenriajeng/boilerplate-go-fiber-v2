package model

import (
	"time"

	"boilerplate-go-fiber-v2/internal/domain/entity"
)

type UserModel struct {
	ID                        uint   `gorm:"primaryKey;autoIncrement"`
	Email                     string `gorm:"uniqueIndex;not null"`
	Username                  string `gorm:"uniqueIndex;not null"`
	Password                  string `gorm:"not null"`
	FirstName                 string `gorm:"not null"`
	LastName                  string `gorm:"not null"`
	Phone                     string `gorm:"uniqueIndex"`
	Avatar                    string
	Role                      string `gorm:"default:'user'"`
	Status                    string `gorm:"default:'active'"`
	EmailVerifiedAt           *time.Time
	PhoneVerifiedAt           *time.Time
	LastLoginAt               *time.Time
	TFAEnabled                bool `gorm:"default:false"`
	TFASecret                 *string
	TFABackupCodes            []string `gorm:"type:text[]"`
	EmailVerificationToken    *string
	EmailVerificationSentAt   *time.Time
	EmailVerificationAttempts int `gorm:"default:0"`
	CreatedAt                 time.Time
	UpdatedAt                 time.Time
}

func (UserModel) TableName() string {
	return "users"
}

func (m *UserModel) ToEntity() *entity.User {
	return &entity.User{
		ID:                        m.ID,
		Email:                     m.Email,
		Username:                  m.Username,
		Password:                  m.Password,
		FirstName:                 m.FirstName,
		LastName:                  m.LastName,
		Phone:                     m.Phone,
		Avatar:                    m.Avatar,
		Role:                      m.Role,
		Status:                    m.Status,
		EmailVerifiedAt:           m.EmailVerifiedAt,
		PhoneVerifiedAt:           m.PhoneVerifiedAt,
		LastLoginAt:               m.LastLoginAt,
		TFAEnabled:                m.TFAEnabled,
		TFASecret:                 m.TFASecret,
		TFABackupCodes:            m.TFABackupCodes,
		EmailVerificationToken:    m.EmailVerificationToken,
		EmailVerificationSentAt:   m.EmailVerificationSentAt,
		EmailVerificationAttempts: m.EmailVerificationAttempts,
		CreatedAt:                 m.CreatedAt,
		UpdatedAt:                 m.UpdatedAt,
	}
}

func (m *UserModel) FromEntity(user *entity.User) {
	m.ID = user.ID
	m.Email = user.Email
	m.Username = user.Username
	m.Password = user.Password
	m.FirstName = user.FirstName
	m.LastName = user.LastName
	m.Phone = user.Phone
	m.Avatar = user.Avatar
	m.Role = user.Role
	m.Status = user.Status
	m.EmailVerifiedAt = user.EmailVerifiedAt
	m.PhoneVerifiedAt = user.PhoneVerifiedAt
	m.LastLoginAt = user.LastLoginAt
	m.TFAEnabled = user.TFAEnabled
	m.TFASecret = user.TFASecret
	m.TFABackupCodes = user.TFABackupCodes
	m.EmailVerificationToken = user.EmailVerificationToken
	m.EmailVerificationSentAt = user.EmailVerificationSentAt
	m.EmailVerificationAttempts = user.EmailVerificationAttempts
	m.CreatedAt = user.CreatedAt
	m.UpdatedAt = user.UpdatedAt
}
