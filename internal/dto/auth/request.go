package auth

type RegisterRequest struct {
	Email     string `json:"email" validate:"required,email"`
	Username  string `json:"username" validate:"required,min=3,max=20"`
	Password  string `json:"password" validate:"required,min=8"`
	FirstName string `json:"first_name" validate:"required"`
	LastName  string `json:"last_name" validate:"required"`
	Phone     string `json:"phone" validate:"required"`
}

type LoginRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

type RefreshTokenRequest struct {
	RefreshToken string `json:"refresh_token" validate:"required"`
}

type PasswordResetRequest struct {
	Email string `json:"email" validate:"required,email"`
}

type ResetPasswordRequest struct {
	Token       string `json:"token" validate:"required"`
	NewPassword string `json:"new_password" validate:"required,min=8"`
}

type TFACodeRequest struct {
	Code string `json:"code" validate:"required,min=6,max=6"`
}

type EnableTFARequest struct {
	Password string `json:"password" validate:"required"`
}

type DisableTFARequest struct {
	Password string `json:"password" validate:"required"`
}

type VerifyTFARequest struct {
	Code string `json:"code" validate:"required,min=6,max=6"`
}
