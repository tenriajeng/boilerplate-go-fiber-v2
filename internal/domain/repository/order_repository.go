package repository

import (
	"boilerplate-go-fiber-v2/internal/domain/entity"
	"context"
)

type OrderRepository interface {
	Create(ctx context.Context, order *entity.Order) error
	GetByID(ctx context.Context, id uint) (*entity.Order, error)
	GetByOrderNumber(ctx context.Context, orderNumber string) (*entity.Order, error)
	GetByUserID(ctx context.Context, userID uint, filter OrderFilter) ([]*entity.Order, error)
	Update(ctx context.Context, order *entity.Order) error
	UpdateStatus(ctx context.Context, id uint, status string) error
	Delete(ctx context.Context, id uint) error
	List(ctx context.Context, filter OrderFilter) ([]*entity.Order, error)
	Count(ctx context.Context, filter OrderFilter) (int64, error)
}

type OrderFilter struct {
	UserID    uint    `json:"user_id"`
	Status    string  `json:"status"`
	MinAmount float64 `json:"min_amount"`
	MaxAmount float64 `json:"max_amount"`
	Page      int     `json:"page"`
	Limit     int     `json:"limit"`
	SortBy    string  `json:"sort_by"`
	SortDesc  bool    `json:"sort_desc"`
}
