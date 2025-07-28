package repository

import (
	"boilerplate-go-fiber-v2/internal/domain/entity"
	"context"
)

type PaymentRepository interface {
	Create(ctx context.Context, payment *entity.Payment) error
	GetByID(ctx context.Context, id uint) (*entity.Payment, error)
	GetByGatewayRef(ctx context.Context, gatewayRef string) (*entity.Payment, error)
	GetByOrderID(ctx context.Context, orderID uint) ([]*entity.Payment, error)
	GetByUserID(ctx context.Context, userID uint, filter PaymentFilter) ([]*entity.Payment, error)
	Update(ctx context.Context, payment *entity.Payment) error
	UpdateStatus(ctx context.Context, id uint, status string) error
	Delete(ctx context.Context, id uint) error
	Count(ctx context.Context, filter PaymentFilter) (int64, error)
	CleanExpiredPayments(ctx context.Context) error
}

type PaymentFilter struct {
	UserID        uint    `json:"user_id"`
	OrderID       uint    `json:"order_id"`
	Status        string  `json:"status"`
	Gateway       string  `json:"gateway"`
	PaymentMethod string  `json:"payment_method"`
	MinAmount     float64 `json:"min_amount"`
	MaxAmount     float64 `json:"max_amount"`
	Page          int     `json:"page"`
	Limit         int     `json:"limit"`
	SortBy        string  `json:"sort_by"`
	SortDesc      bool    `json:"sort_desc"`
}
