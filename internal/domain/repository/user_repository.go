package repository

import (
	"boilerplate-go-fiber-v2/internal/domain/entity"
	"context"
)

type UserRepository interface {
	Create(ctx context.Context, user *entity.User) error
	GetByID(ctx context.Context, id uint) (*entity.User, error)
	GetByEmail(ctx context.Context, email string) (*entity.User, error)
	GetByUsername(ctx context.Context, username string) (*entity.User, error)
	Update(ctx context.Context, user *entity.User) error
	Delete(ctx context.Context, id uint) error
	List(ctx context.Context, filter UserFilter) ([]*entity.User, error)
	Count(ctx context.Context, filter UserFilter) (int64, error)
	UpdateLastLogin(ctx context.Context, userID uint) error
	UpdateStatus(ctx context.Context, userID uint, status string) error
	UpdateTFA(ctx context.Context, userID uint, enabled bool, secret string, backupCodes []string) error
}

type UserFilter struct {
	Search   string `json:"search"`
	Role     string `json:"role"`
	Status   string `json:"status"`
	Page     int    `json:"page"`
	Limit    int    `json:"limit"`
	SortBy   string `json:"sort_by"`
	SortDesc bool   `json:"sort_desc"`
}
