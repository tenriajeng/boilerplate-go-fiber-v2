package repository

import (
	"context"
	"errors"

	"boilerplate-go-fiber-v2/internal/domain/entity"
	"boilerplate-go-fiber-v2/internal/domain/repository"

	"gorm.io/gorm"
)

type orderRepository struct {
	db *gorm.DB
}

// NewOrderRepository creates a new order repository
func NewOrderRepository(db *gorm.DB) repository.OrderRepository {
	return &orderRepository{db: db}
}

// Create creates a new order
func (r *orderRepository) Create(ctx context.Context, order *entity.Order) error {
	return r.db.WithContext(ctx).Create(order).Error
}

// GetByID gets an order by ID
func (r *orderRepository) GetByID(ctx context.Context, id uint) (*entity.Order, error) {
	var order entity.Order
	err := r.db.WithContext(ctx).First(&order, id).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("order not found")
		}
		return nil, err
	}
	return &order, nil
}

// GetByOrderNumber gets an order by order number
func (r *orderRepository) GetByOrderNumber(ctx context.Context, orderNumber string) (*entity.Order, error) {
	var order entity.Order
	err := r.db.WithContext(ctx).Where("order_number = ?", orderNumber).First(&order).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("order not found")
		}
		return nil, err
	}
	return &order, nil
}

// GetByUserID gets orders by user ID with filtering
func (r *orderRepository) GetByUserID(ctx context.Context, userID uint, filter repository.OrderFilter) ([]*entity.Order, error) {
	var orders []*entity.Order
	query := r.db.WithContext(ctx).Where("user_id = ?", userID)

	// Apply filters
	if filter.Status != "" {
		query = query.Where("status = ?", filter.Status)
	}

	if filter.MinAmount > 0 {
		query = query.Where("total_amount >= ?", filter.MinAmount)
	}

	if filter.MaxAmount > 0 {
		query = query.Where("total_amount <= ?", filter.MaxAmount)
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

	err := query.Find(&orders).Error
	return orders, err
}

// Update updates an order
func (r *orderRepository) Update(ctx context.Context, order *entity.Order) error {
	return r.db.WithContext(ctx).Save(order).Error
}

// UpdateStatus updates order status
func (r *orderRepository) UpdateStatus(ctx context.Context, id uint, status string) error {
	return r.db.WithContext(ctx).Model(&entity.Order{}).Where("id = ?", id).Update("status", status).Error
}

// Delete deletes an order
func (r *orderRepository) Delete(ctx context.Context, id uint) error {
	return r.db.WithContext(ctx).Delete(&entity.Order{}, id).Error
}

// List gets orders with filtering and pagination
func (r *orderRepository) List(ctx context.Context, filter repository.OrderFilter) ([]*entity.Order, error) {
	var orders []*entity.Order
	query := r.db.WithContext(ctx)

	// Apply filters
	if filter.UserID > 0 {
		query = query.Where("user_id = ?", filter.UserID)
	}

	if filter.Status != "" {
		query = query.Where("status = ?", filter.Status)
	}

	if filter.MinAmount > 0 {
		query = query.Where("total_amount >= ?", filter.MinAmount)
	}

	if filter.MaxAmount > 0 {
		query = query.Where("total_amount <= ?", filter.MaxAmount)
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

	err := query.Find(&orders).Error
	return orders, err
}

// Count counts orders with filtering
func (r *orderRepository) Count(ctx context.Context, filter repository.OrderFilter) (int64, error) {
	var count int64
	query := r.db.WithContext(ctx).Model(&entity.Order{})

	// Apply filters
	if filter.UserID > 0 {
		query = query.Where("user_id = ?", filter.UserID)
	}

	if filter.Status != "" {
		query = query.Where("status = ?", filter.Status)
	}

	if filter.MinAmount > 0 {
		query = query.Where("total_amount >= ?", filter.MinAmount)
	}

	if filter.MaxAmount > 0 {
		query = query.Where("total_amount <= ?", filter.MaxAmount)
	}

	err := query.Count(&count).Error
	return count, err
}
