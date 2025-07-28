package handler

import (
	"boilerplate-go-fiber-v2/internal/domain/entity"
	"boilerplate-go-fiber-v2/internal/domain/service"
	"boilerplate-go-fiber-v2/internal/dto/auth"
	"boilerplate-go-fiber-v2/pkg/response"
	"boilerplate-go-fiber-v2/pkg/utils"
	"boilerplate-go-fiber-v2/pkg/validator"

	"github.com/gofiber/fiber/v2"
)

type AuthHandler struct {
	authService service.AuthService
	userService service.UserService
}

// NewAuthHandler creates a new auth handler
func NewAuthHandler(authService service.AuthService, userService service.UserService) *AuthHandler {
	return &AuthHandler{
		authService: authService,
		userService: userService,
	}
}

// Register handles user registration
func (h *AuthHandler) Register(c *fiber.Ctx) error {
	var req auth.RegisterRequest
	if err := c.BodyParser(&req); err != nil {
		return response.ValidationError(c, "Invalid request body")
	}

	// Validate request
	if err := validator.ValidateStruct(req); err != nil {
		return response.ValidationError(c, err.Error())
	}

	// Create user entity
	user := &entity.User{
		Email:     req.Email,
		Username:  req.Username,
		Password:  req.Password,
		FirstName: req.FirstName,
		LastName:  req.LastName,
		Phone:     req.Phone,
	}

	// Register user
	err := h.authService.Register(c.Context(), user)
	if err != nil {
		return response.Error(c, err.Error(), fiber.StatusBadRequest)
	}

	// Create response
	resp := auth.RegisterResponse{
		User:    h.mapUserToResponse(user),
		Message: "User registered successfully",
	}

	return response.Success(c, "Registration successful", resp)
}

// Login handles user login
func (h *AuthHandler) Login(c *fiber.Ctx) error {
	var req auth.LoginRequest
	if err := c.BodyParser(&req); err != nil {
		return response.ValidationError(c, "Invalid request body")
	}

	// Validate request
	if err := validator.ValidateStruct(req); err != nil {
		return response.ValidationError(c, err.Error())
	}

	// Login user
	user, session, err := h.authService.Login(c.Context(), req.Email, req.Password)
	if err != nil {
		return response.Unauthorized(c, err.Error())
	}

	// Create response
	resp := auth.LoginResponse{
		User:         h.mapUserToResponse(user),
		AccessToken:  session.Token,
		RefreshToken: session.RefreshToken,
		ExpiresAt:    session.ExpiresAt,
		TokenType:    "Bearer",
	}

	return response.Success(c, "Login successful", resp)
}

// Logout handles user logout
func (h *AuthHandler) Logout(c *fiber.Ctx) error {
	token := c.Get("Authorization")
	if token == "" {
		return response.Unauthorized(c, "Token required")
	}

	// Remove "Bearer " prefix
	if len(token) > 7 && token[:7] == "Bearer " {
		token = token[7:]
	}

	err := h.authService.Logout(c.Context(), token)
	if err != nil {
		return response.Error(c, err.Error(), fiber.StatusBadRequest)
	}

	resp := auth.LogoutResponse{
		Message: "Logged out successfully",
	}

	return response.Success(c, "Logout successful", resp)
}

// RefreshToken handles token refresh
func (h *AuthHandler) RefreshToken(c *fiber.Ctx) error {
	var req auth.RefreshTokenRequest
	if err := c.BodyParser(&req); err != nil {
		return response.ValidationError(c, "Invalid request body")
	}

	// Validate request
	if err := validator.ValidateStruct(req); err != nil {
		return response.ValidationError(c, err.Error())
	}

	// Refresh token
	session, err := h.authService.RefreshToken(c.Context(), req.RefreshToken)
	if err != nil {
		return response.Unauthorized(c, err.Error())
	}

	// Create response
	resp := auth.RefreshTokenResponse{
		AccessToken:  session.Token,
		RefreshToken: session.RefreshToken,
		ExpiresAt:    session.ExpiresAt,
		TokenType:    "Bearer",
	}

	return response.Success(c, "Token refreshed successfully", resp)
}

// CreatePasswordReset handles password reset request
func (h *AuthHandler) CreatePasswordReset(c *fiber.Ctx) error {
	var req auth.PasswordResetRequest
	if err := c.BodyParser(&req); err != nil {
		return response.ValidationError(c, "Invalid request body")
	}

	// Validate request
	if err := validator.ValidateStruct(req); err != nil {
		return response.ValidationError(c, err.Error())
	}

	// Create password reset
	err := h.authService.CreatePasswordReset(c.Context(), req.Email)
	if err != nil {
		return response.Error(c, err.Error(), fiber.StatusBadRequest)
	}

	resp := auth.PasswordResetResponse{
		Message: "Password reset email sent",
	}

	return response.Success(c, "Password reset initiated", resp)
}

// ResetPassword handles password reset
func (h *AuthHandler) ResetPassword(c *fiber.Ctx) error {
	var req auth.ResetPasswordRequest
	if err := c.BodyParser(&req); err != nil {
		return response.ValidationError(c, "Invalid request body")
	}

	// Validate request
	if err := validator.ValidateStruct(req); err != nil {
		return response.ValidationError(c, err.Error())
	}

	// Reset password
	err := h.authService.ResetPassword(c.Context(), req.Token, req.NewPassword)
	if err != nil {
		return response.Error(c, err.Error(), fiber.StatusBadRequest)
	}

	resp := auth.PasswordResetResponse{
		Message: "Password reset successfully",
	}

	return response.Success(c, "Password reset successful", resp)
}

// CreateTFACode creates TFA code
func (h *AuthHandler) CreateTFACode(c *fiber.Ctx) error {
	userID := c.Locals("user_id").(uint)

	err := h.authService.CreateTFACode(c.Context(), userID)
	if err != nil {
		return response.Error(c, err.Error(), fiber.StatusBadRequest)
	}

	resp := auth.TFACodeResponse{
		Message: "TFA code sent",
	}

	return response.Success(c, "TFA code created", resp)
}

// EnableTFA enables TFA for user
func (h *AuthHandler) EnableTFA(c *fiber.Ctx) error {
	var req auth.EnableTFARequest
	if err := c.BodyParser(&req); err != nil {
		return response.ValidationError(c, "Invalid request body")
	}

	// Validate request
	if err := validator.ValidateStruct(req); err != nil {
		return response.ValidationError(c, err.Error())
	}

	userID := c.Locals("user_id").(uint)

	// Verify password first
	user, err := h.userService.GetByID(c.Context(), userID)
	if err != nil {
		return response.Error(c, "User not found", fiber.StatusNotFound)
	}

	if !utils.CheckPassword(req.Password, user.Password) {
		return response.Unauthorized(c, "Invalid password")
	}

	// Enable TFA
	user, err = h.authService.EnableTFA(c.Context(), userID)
	if err != nil {
		return response.Error(c, err.Error(), fiber.StatusBadRequest)
	}

	resp := auth.TFAResponse{
		Secret:      utils.SafePtr(user.TFASecret, ""),
		QRCode:      "", // TODO: Generate QR code
		BackupCodes: user.TFABackupCodes,
		Message:     "TFA enabled successfully",
	}

	return response.Success(c, "TFA enabled", resp)
}

// DisableTFA disables TFA for user
func (h *AuthHandler) DisableTFA(c *fiber.Ctx) error {
	var req auth.DisableTFARequest
	if err := c.BodyParser(&req); err != nil {
		return response.ValidationError(c, "Invalid request body")
	}

	// Validate request
	if err := validator.ValidateStruct(req); err != nil {
		return response.ValidationError(c, err.Error())
	}

	userID := c.Locals("user_id").(uint)

	// Verify password first
	user, err := h.userService.GetByID(c.Context(), userID)
	if err != nil {
		return response.Error(c, "User not found", fiber.StatusNotFound)
	}

	if !utils.CheckPassword(req.Password, user.Password) {
		return response.Unauthorized(c, "Invalid password")
	}

	// Disable TFA
	err = h.authService.DisableTFA(c.Context(), userID)
	if err != nil {
		return response.Error(c, err.Error(), fiber.StatusBadRequest)
	}

	resp := auth.TFADisableResponse{
		Message: "TFA disabled successfully",
	}

	return response.Success(c, "TFA disabled", resp)
}

// VerifyTFA verifies TFA code
func (h *AuthHandler) VerifyTFA(c *fiber.Ctx) error {
	var req auth.VerifyTFARequest
	if err := c.BodyParser(&req); err != nil {
		return response.ValidationError(c, "Invalid request body")
	}

	// Validate request
	if err := validator.ValidateStruct(req); err != nil {
		return response.ValidationError(c, err.Error())
	}

	userID := c.Locals("user_id").(uint)

	// Verify TFA
	err := h.authService.VerifyTFA(c.Context(), userID, req.Code)
	if err != nil {
		return response.Unauthorized(c, err.Error())
	}

	resp := auth.TFACodeResponse{
		Message: "TFA verified successfully",
	}

	return response.Success(c, "TFA verified", resp)
}

// Helper method to map user entity to response
func (h *AuthHandler) mapUserToResponse(user *entity.User) auth.UserResponse {
	return auth.UserResponse{
		ID:              user.ID,
		Email:           user.Email,
		Username:        user.Username,
		FirstName:       user.FirstName,
		LastName:        user.LastName,
		Phone:           user.Phone,
		Avatar:          user.Avatar,
		Role:            user.Role,
		Status:          user.Status,
		EmailVerifiedAt: user.EmailVerifiedAt,
		PhoneVerifiedAt: user.PhoneVerifiedAt,
		LastLoginAt:     user.LastLoginAt,
		TFAEnabled:      user.TFAEnabled,
		CreatedAt:       user.CreatedAt,
		UpdatedAt:       user.UpdatedAt,
	}
}
