package repository

import (
	"context"
	"errors"
	"time"

	"boilerplate-go-fiber-v2/internal/domain/entity"
	"boilerplate-go-fiber-v2/internal/domain/repository"

	"gorm.io/gorm"
)

type paymentRepository struct {
	db *gorm.DB
}

// NewPaymentRepository creates a new payment repository
func NewPaymentRepository(db *gorm.DB) repository.PaymentRepository {
	return &paymentRepository{db: db}
}

// Create creates a new payment
func (r *paymentRepository) Create(ctx context.Context, payment *entity.Payment) error {
	return r.db.WithContext(ctx).Create(payment).Error
}

// GetByID gets a payment by ID
func (r *paymentRepository) GetByID(ctx context.Context, id uint) (*entity.Payment, error) {
	var payment entity.Payment
	err := r.db.WithContext(ctx).First(&payment, id).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("payment not found")
		}
		return nil, err
	}
	return &payment, nil
}

// GetByGatewayRef gets a payment by gateway reference
func (r *paymentRepository) GetByGatewayRef(ctx context.Context, gatewayRef string) (*entity.Payment, error) {
	var payment entity.Payment
	err := r.db.WithContext(ctx).Where("gateway_ref = ?", gatewayRef).First(&payment).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("payment not found")
		}
		return nil, err
	}
	return &payment, nil
}

// GetByOrderID gets payments by order ID
func (r *paymentRepository) GetByOrderID(ctx context.Context, orderID uint) ([]*entity.Payment, error) {
	var payments []*entity.Payment
	err := r.db.WithContext(ctx).Where("order_id = ?", orderID).Find(&payments).Error
	return payments, err
}

// GetByUserID gets payments by user ID with filtering
func (r *paymentRepository) GetByUserID(ctx context.Context, userID uint, filter repository.PaymentFilter) ([]*entity.Payment, error) {
	var payments []*entity.Payment
	query := r.db.WithContext(ctx).Where("user_id = ?", userID)

	// Apply filters
	if filter.OrderID > 0 {
		query = query.Where("order_id = ?", filter.OrderID)
	}

	if filter.Status != "" {
		query = query.Where("status = ?", filter.Status)
	}

	if filter.Gateway != "" {
		query = query.Where("gateway = ?", filter.Gateway)
	}

	if filter.PaymentMethod != "" {
		query = query.Where("payment_method = ?", filter.PaymentMethod)
	}

	if filter.MinAmount > 0 {
		query = query.Where("amount >= ?", filter.MinAmount)
	}

	if filter.MaxAmount > 0 {
		query = query.Where("amount <= ?", filter.MaxAmount)
	}

	// Apply sorting
	sortBy := filter.SortBy
	if sortBy == "" {
		sortBy = "created_at"
	}
	if filter.SortDesc {
		sortBy += " DESC"
	}
	query = query.Order(sortBy)

	// Apply pagination
	if filter.Page > 0 && filter.Limit > 0 {
		offset := (filter.Page - 1) * filter.Limit
		query = query.Offset(offset).Limit(filter.Limit)
	}

	err := query.Find(&payments).Error
	return payments, err
}

// Update updates a payment
func (r *paymentRepository) Update(ctx context.Context, payment *entity.Payment) error {
	return r.db.WithContext(ctx).Save(payment).Error
}

// UpdateStatus updates payment status
func (r *paymentRepository) UpdateStatus(ctx context.Context, id uint, status string) error {
	return r.db.WithContext(ctx).Model(&entity.Payment{}).Where("id = ?", id).Update("status", status).Error
}

// Delete deletes a payment
func (r *paymentRepository) Delete(ctx context.Context, id uint) error {
	return r.db.WithContext(ctx).Delete(&entity.Payment{}, id).Error
}

// Count counts payments with filtering
func (r *paymentRepository) Count(ctx context.Context, filter repository.PaymentFilter) (int64, error) {
	var count int64
	query := r.db.WithContext(ctx).Model(&entity.Payment{})

	// Apply filters
	if filter.UserID > 0 {
		query = query.Where("user_id = ?", filter.UserID)
	}

	if filter.OrderID > 0 {
		query = query.Where("order_id = ?", filter.OrderID)
	}

	if filter.Status != "" {
		query = query.Where("status = ?", filter.Status)
	}

	if filter.Gateway != "" {
		query = query.Where("gateway = ?", filter.Gateway)
	}

	if filter.PaymentMethod != "" {
		query = query.Where("payment_method = ?", filter.PaymentMethod)
	}

	if filter.MinAmount > 0 {
		query = query.Where("amount >= ?", filter.MinAmount)
	}

	if filter.MaxAmount > 0 {
		query = query.Where("amount <= ?", filter.MaxAmount)
	}

	err := query.Count(&count).Error
	return count, err
}

// CleanExpiredPayments removes expired payments
func (r *paymentRepository) CleanExpiredPayments(ctx context.Context) error {
	return r.db.WithContext(ctx).Where("expires_at < ? AND status = ?", time.Now(), "pending").Delete(&entity.Payment{}).Error
}
