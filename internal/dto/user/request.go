package user

type UpdateProfileRequest struct {
	FirstName string `json:"first_name" validate:"required"`
	LastName  string `json:"last_name" validate:"required"`
	Phone     string `json:"phone" validate:"required"`
	Avatar    string `json:"avatar"`
}

type ChangePasswordRequest struct {
	OldPassword string `json:"old_password" validate:"required"`
	NewPassword string `json:"new_password" validate:"required,min=8"`
}

type UpdateStatusRequest struct {
	Status string `json:"status" validate:"required,oneof=active inactive suspended banned"`
}

type UserFilterRequest struct {
	Search   string `query:"search"`
	Role     string `query:"role"`
	Status   string `query:"status"`
	Page     int    `query:"page" validate:"min=1"`
	Limit    int    `query:"limit" validate:"min=1,max=100"`
	SortBy   string `query:"sort_by"`
	SortDesc bool   `query:"sort_desc"`
}
