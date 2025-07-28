package entity

import (
	"time"
)

type User struct {
	ID                        uint
	Email                     string
	Username                  string
	Password                  string
	FirstName                 string
	LastName                  string
	Phone                     string
	Avatar                    string
	Role                      string
	Status                    string
	EmailVerifiedAt           *time.Time
	PhoneVerifiedAt           *time.Time
	LastLoginAt               *time.Time
	TFAEnabled                bool
	TFASecret                 *string
	TFABackupCodes            []string
	EmailVerificationToken    *string
	EmailVerificationSentAt   *time.Time
	EmailVerificationAttempts int
	CreatedAt                 time.Time
	UpdatedAt                 time.Time
}

// IsActive checks if user is active
func (u *User) IsActive() bool {
	return u.Status == "active"
}

// IsEmailVerified checks if user email is verified
func (u *User) IsEmailVerified() bool {
	return u.EmailVerifiedAt != nil
}

// IsPhoneVerified checks if user phone is verified
func (u *User) IsPhoneVerified() bool {
	return u.PhoneVerifiedAt != nil
}

// IsAdmin checks if user is admin
func (u *User) IsAdmin() bool {
	return u.Role == "admin"
}

// IsTFAEnabled checks if TFA is enabled
func (u *User) IsTFAEnabled() bool {
	return u.TFAEnabled
}

// GetFullName returns user's full name
func (u *User) GetFullName() string {
	if u.FirstName != "" && u.LastName != "" {
		return u.FirstName + " " + u.LastName
	}
	if u.FirstName != "" {
		return u.FirstName
	}
	if u.LastName != "" {
		return u.LastName
	}
	return u.Username
}

// CanVerifyEmail checks if user can request email verification
func (u *User) CanVerifyEmail() bool {
	return !u.IsEmailVerified() && u.EmailVerificationAttempts < 5
}

// HasValidEmailVerificationToken checks if user has a valid email verification token
func (u *User) HasValidEmailVerificationToken() bool {
	return u.EmailVerificationToken != nil &&
		u.EmailVerificationSentAt != nil &&
		time.Since(*u.EmailVerificationSentAt) < 24*time.Hour
}

// EnableTFA enables TFA for user
func (u *User) EnableTFA(secret string, backupCodes []string) {
	u.TFAEnabled = true
	u.TFASecret = &secret
	u.TFABackupCodes = backupCodes
}

// DisableTFA disables TFA for user
func (u *User) DisableTFA() {
	u.TFAEnabled = false
	u.TFASecret = nil
	u.TFABackupCodes = nil
}

// MarkEmailVerified marks email as verified
func (u *User) MarkEmailVerified() {
	now := time.Now()
	u.EmailVerifiedAt = &now
}
