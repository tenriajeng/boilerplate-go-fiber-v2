package entity

import (
	"time"
)

type AuthSession struct {
	ID           uint
	UserID       uint
	Token        string
	RefreshToken string
	ExpiresAt    time.Time
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

type PasswordReset struct {
	ID        uint
	UserID    uint
	Token     string
	ExpiresAt time.Time
	Used      bool
	CreatedAt time.Time
	UpdatedAt time.Time
}

type TFACode struct {
	ID        uint
	UserID    uint
	Code      string
	ExpiresAt time.Time
	Used      bool
	CreatedAt time.Time
	UpdatedAt time.Time
}

type EmailVerification struct {
	ID         uint
	UserID     uint
	Email      string
	Token      string
	ExpiresAt  time.Time
	VerifiedAt *time.Time
	CreatedAt  time.Time
	UpdatedAt  time.Time
}

// Business methods for AuthSession
func (a *AuthSession) IsExpired() bool {
	return time.Now().After(a.ExpiresAt)
}

func (a *AuthSession) IsValid() bool {
	return !a.IsExpired()
}

// Business methods for PasswordReset
func (p *PasswordReset) IsExpired() bool {
	return time.Now().After(p.ExpiresAt)
}

func (p *PasswordReset) IsValid() bool {
	return !p.IsExpired() && !p.Used
}

func (p *PasswordReset) MarkAsUsed() {
	p.Used = true
}

// Business methods for TFACode
func (t *TFACode) IsExpired() bool {
	return time.Now().After(t.ExpiresAt)
}

func (t *TFACode) IsValid() bool {
	return !t.IsExpired() && !t.Used
}

func (t *TFACode) MarkAsUsed() {
	t.Used = true
}

// Business methods for EmailVerification
func (ev *EmailVerification) IsExpired() bool {
	return time.Now().After(ev.ExpiresAt)
}

func (ev *EmailVerification) IsValid() bool {
	return ev.VerifiedAt == nil && time.Now().Before(ev.ExpiresAt)
}

func (ev *EmailVerification) MarkVerified() {
	now := time.Now()
	ev.VerifiedAt = &now
}
