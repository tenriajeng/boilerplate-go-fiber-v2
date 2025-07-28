package user

import "time"

type UserResponse struct {
	ID              uint       `json:"id"`
	Email           string     `json:"email"`
	Username        string     `json:"username"`
	FirstName       string     `json:"first_name"`
	LastName        string     `json:"last_name"`
	Phone           string     `json:"phone"`
	Avatar          string     `json:"avatar"`
	Role            string     `json:"role"`
	Status          string     `json:"status"`
	EmailVerifiedAt *time.Time `json:"email_verified_at"`
	PhoneVerifiedAt *time.Time `json:"phone_verified_at"`
	LastLoginAt     *time.Time `json:"last_login_at"`
	TFAEnabled      bool       `json:"tfa_enabled"`
	CreatedAt       time.Time  `json:"created_at"`
	UpdatedAt       time.Time  `json:"updated_at"`
}

type UserListResponse struct {
	Users []UserResponse `json:"users"`
	Meta  MetaResponse   `json:"meta"`
}

type MetaResponse struct {
	Page       int   `json:"page"`
	Limit      int   `json:"limit"`
	Total      int64 `json:"total"`
	TotalPages int   `json:"total_pages"`
}

type UpdateProfileResponse struct {
	User    UserResponse `json:"user"`
	Message string       `json:"message"`
}

type ChangePasswordResponse struct {
	Message string `json:"message"`
}

type UpdateStatusResponse struct {
	User    UserResponse `json:"user"`
	Message string       `json:"message"`
}
