package service

import (
	"context"
	"errors"
	"time"

	"boilerplate-go-fiber-v2/config"
	"boilerplate-go-fiber-v2/internal/domain/entity"
	"boilerplate-go-fiber-v2/internal/domain/repository"
	"boilerplate-go-fiber-v2/pkg/jwt"
	"boilerplate-go-fiber-v2/pkg/utils"
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

type authService struct {
	userRepo    repository.UserRepository
	authRepo    repository.AuthRepository
	userService UserService
	config      *config.Config
}

// NewAuthService creates a new auth service
func NewAuthService(userRepo repository.UserRepository, authRepo repository.AuthRepository, userService UserService, config *config.Config) AuthService {
	return &authService{
		userRepo:    userRepo,
		authRepo:    authRepo,
		userService: userService,
		config:      config,
	}
}

// Login authenticates a user
func (s *authService) Login(ctx context.Context, email, password string) (*entity.User, *entity.AuthSession, error) {
	// Get user by email
	user, err := s.userRepo.GetByEmail(ctx, email)
	if err != nil {
		return nil, nil, errors.New("invalid credentials")
	}

	// Check if user is active
	if !user.IsActive() {
		return nil, nil, errors.New("account is not active")
	}

	// Check if email is verified
	if !user.IsEmailVerified() {
		return nil, nil, errors.New("email not verified")
	}

	// Verify password
	if !utils.CheckPassword(password, user.Password) {
		return nil, nil, errors.New("invalid credentials")
	}

	// Generate tokens
	accessToken, err := jwt.GenerateToken(user.ID, user.Email, user.Role, s.config.JWT.Secret, s.config.JWT.Expiry)
	if err != nil {
		return nil, nil, err
	}

	refreshToken, err := utils.GenerateSecureToken(32)
	if err != nil {
		return nil, nil, err
	}

	// Create session
	session := &entity.AuthSession{
		UserID:       user.ID,
		Token:        accessToken,
		RefreshToken: refreshToken,
		ExpiresAt:    time.Now().Add(s.config.JWT.Expiry),
		CreatedAt:    time.Now(),
	}

	err = s.authRepo.CreateSession(ctx, session)
	if err != nil {
		return nil, nil, err
	}

	// Update last login
	err = s.userService.UpdateLastLogin(ctx, user.ID)
	if err != nil {
		// Log error but don't fail login
	}

	return user, session, nil
}

// Logout logs out a user
func (s *authService) Logout(ctx context.Context, token string) error {
	return s.authRepo.DeleteSession(ctx, token)
}

// RefreshToken refreshes an access token
func (s *authService) RefreshToken(ctx context.Context, refreshToken string) (*entity.AuthSession, error) {
	// Get session by refresh token
	session, err := s.authRepo.GetSessionByRefreshToken(ctx, refreshToken)
	if err != nil {
		return nil, errors.New("invalid refresh token")
	}

	// Check if session is expired
	if session.IsExpired() {
		return nil, errors.New("session expired")
	}

	// Get user
	user, err := s.userRepo.GetByID(ctx, session.UserID)
	if err != nil {
		return nil, errors.New("user not found")
	}

	// Generate new tokens
	accessToken, err := jwt.GenerateToken(user.ID, user.Email, user.Role, s.config.JWT.Secret, s.config.JWT.Expiry)
	if err != nil {
		return nil, err
	}

	newRefreshToken, err := utils.GenerateSecureToken(32)
	if err != nil {
		return nil, err
	}

	// Update session
	session.Token = accessToken
	session.RefreshToken = newRefreshToken
	session.ExpiresAt = time.Now().Add(s.config.JWT.Expiry)
	session.CreatedAt = time.Now()

	err = s.authRepo.UpdateSession(ctx, session)
	if err != nil {
		return nil, err
	}

	return session, nil
}

// Register registers a new user
func (s *authService) Register(ctx context.Context, user *entity.User) error {
	return s.userService.Register(ctx, user)
}

// ValidateToken validates a JWT token
func (s *authService) ValidateToken(ctx context.Context, token string) (*jwt.Claims, error) {
	claims, err := jwt.ValidateToken(token, s.config.JWT.Secret)
	if err != nil {
		return nil, err
	}

	// Check if session exists
	_, err = s.authRepo.GetSessionByToken(ctx, token)
	if err != nil {
		return nil, errors.New("session not found")
	}

	return claims, nil
}

// CreatePasswordReset creates a password reset request
func (s *authService) CreatePasswordReset(ctx context.Context, email string) error {
	user, err := s.userRepo.GetByEmail(ctx, email)
	if err != nil {
		return errors.New("user not found")
	}

	// Generate reset token
	token, err := utils.GenerateSecureToken(32)
	if err != nil {
		return err
	}

	// Create password reset
	reset := &entity.PasswordReset{
		UserID:    user.ID,
		Token:     token,
		ExpiresAt: time.Now().Add(24 * time.Hour), // 24 hours
		CreatedAt: time.Now(),
	}

	return s.authRepo.CreatePasswordReset(ctx, reset)
}

// ResetPassword resets user password
func (s *authService) ResetPassword(ctx context.Context, token, newPassword string) error {
	// Get password reset
	reset, err := s.authRepo.GetPasswordResetByToken(ctx, token)
	if err != nil {
		return errors.New("invalid reset token")
	}

	// Check if reset is valid
	if !reset.IsValid() {
		return errors.New("reset token expired or already used")
	}

	// Hash new password
	hashedPassword, err := utils.HashPassword(newPassword)
	if err != nil {
		return err
	}

	// Update user password
	user, err := s.userRepo.GetByID(ctx, reset.UserID)
	if err != nil {
		return errors.New("user not found")
	}

	user.Password = hashedPassword
	user.UpdatedAt = time.Now()

	err = s.userRepo.Update(ctx, user)
	if err != nil {
		return err
	}

	// Mark reset as used
	return s.authRepo.MarkPasswordResetUsed(ctx, token)
}

// CreateTFACode creates a TFA code for user
func (s *authService) CreateTFACode(ctx context.Context, userID uint) error {
	// Generate TFA code
	code := utils.GenerateTFACode()

	// Create TFA code
	tfaCode := &entity.TFACode{
		UserID:    userID,
		Code:      code,
		ExpiresAt: time.Now().Add(5 * time.Minute), // 5 minutes
		CreatedAt: time.Now(),
	}

	return s.authRepo.CreateTFACode(ctx, tfaCode)
}

// VerifyTFACode verifies a TFA code
func (s *authService) VerifyTFACode(ctx context.Context, userID uint, code string) error {
	tfaCode, err := s.authRepo.GetTFACodeByCode(ctx, code)
	if err != nil {
		return errors.New("invalid TFA code")
	}

	// Check if code belongs to user
	if tfaCode.UserID != userID {
		return errors.New("invalid TFA code")
	}

	// Check if code is valid
	if !tfaCode.IsValid() {
		return errors.New("TFA code expired or already used")
	}

	// Mark code as used
	return s.authRepo.MarkTFACodeUsed(ctx, code)
}

// EnableTFA enables TFA for user
func (s *authService) EnableTFA(ctx context.Context, userID uint) (*entity.User, error) {
	user, err := s.userRepo.GetByID(ctx, userID)
	if err != nil {
		return nil, errors.New("user not found")
	}

	// Generate TFA secret
	secret, err := utils.GenerateSecureToken(32)
	if err != nil {
		return nil, err
	}

	// Generate backup codes
	backupCodes, err := utils.GenerateBackupCodes(8)
	if err != nil {
		return nil, err
	}

	// Enable TFA
	user.EnableTFA(secret, backupCodes)
	user.UpdatedAt = time.Now()

	err = s.userRepo.Update(ctx, user)
	if err != nil {
		return nil, err
	}

	return user, nil
}

// DisableTFA disables TFA for user
func (s *authService) DisableTFA(ctx context.Context, userID uint) error {
	user, err := s.userRepo.GetByID(ctx, userID)
	if err != nil {
		return errors.New("user not found")
	}

	// Disable TFA
	user.DisableTFA()
	user.UpdatedAt = time.Now()

	return s.userRepo.Update(ctx, user)
}

// VerifyTFA verifies TFA for login
func (s *authService) VerifyTFA(ctx context.Context, userID uint, code string) error {
	user, err := s.userRepo.GetByID(ctx, userID)
	if err != nil {
		return errors.New("user not found")
	}

	// Check if TFA is enabled
	if !user.IsTFAEnabled() {
		return errors.New("TFA not enabled")
	}

	// Verify TFA code
	return s.VerifyTFACode(ctx, userID, code)
}
