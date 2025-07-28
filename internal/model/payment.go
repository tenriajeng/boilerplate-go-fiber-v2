package model

import (
	"boilerplate-go-fiber-v2/internal/domain/entity"
	"time"
)

// PaymentModel represents the database model for Payment entity
type PaymentModel struct {
	ID            uint    `gorm:"primaryKey"`
	OrderID       uint    `gorm:"not null"`
	UserID        uint    `gorm:"not null"`
	Amount        float64 `gorm:"not null"`
	Currency      string  `gorm:"default:IDR"`
	PaymentMethod string  `gorm:"not null"`
	Status        string  `gorm:"default:pending"`
	Gateway       string  `gorm:"not null"`
	GatewayRef    string
	GatewayData   entity.JSONB `gorm:"type:jsonb"`
	ExpiresAt     *time.Time
	PaidAt        *time.Time
	CreatedAt     time.Time
	UpdatedAt     time.Time
}

func (PaymentModel) TableName() string {
	return "payments"
}

// ToEntity converts PaymentModel to Payment entity
func (m *PaymentModel) ToEntity() *entity.Payment {
	return &entity.Payment{
		ID:            m.ID,
		OrderID:       m.OrderID,
		UserID:        m.UserID,
		Amount:        m.Amount,
		Currency:      m.Currency,
		PaymentMethod: m.PaymentMethod,
		Status:        m.Status,
		Gateway:       m.Gateway,
		GatewayRef:    m.GatewayRef,
		GatewayData:   m.GatewayData,
		ExpiresAt:     m.ExpiresAt,
		PaidAt:        m.PaidAt,
		CreatedAt:     m.CreatedAt,
		UpdatedAt:     m.UpdatedAt,
	}
}

// FromEntity converts Payment entity to PaymentModel
func (m *PaymentModel) FromEntity(payment *entity.Payment) {
	m.ID = payment.ID
	m.OrderID = payment.OrderID
	m.UserID = payment.UserID
	m.Amount = payment.Amount
	m.Currency = payment.Currency
	m.PaymentMethod = payment.PaymentMethod
	m.Status = payment.Status
	m.Gateway = payment.Gateway
	m.GatewayRef = payment.GatewayRef
	m.GatewayData = payment.GatewayData
	m.ExpiresAt = payment.ExpiresAt
	m.PaidAt = payment.PaidAt
	m.CreatedAt = payment.CreatedAt
	m.UpdatedAt = payment.UpdatedAt
}

// OrderModel represents the database model for Order entity
type OrderModel struct {
	ID          uint    `gorm:"primaryKey"`
	UserID      uint    `gorm:"not null"`
	OrderNumber string  `gorm:"uniqueIndex;not null"`
	TotalAmount float64 `gorm:"not null"`
	Status      string  `gorm:"default:pending"`
	Note        string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

func (OrderModel) TableName() string {
	return "orders"
}

// ToEntity converts OrderModel to Order entity
func (m *OrderModel) ToEntity() *entity.Order {
	return &entity.Order{
		ID:          m.ID,
		UserID:      m.UserID,
		OrderNumber: m.OrderNumber,
		TotalAmount: m.TotalAmount,
		Status:      m.Status,
		Note:        m.Note,
		CreatedAt:   m.CreatedAt,
		UpdatedAt:   m.UpdatedAt,
	}
}

// FromEntity converts Order entity to OrderModel
func (m *OrderModel) FromEntity(order *entity.Order) {
	m.ID = order.ID
	m.UserID = order.UserID
	m.OrderNumber = order.OrderNumber
	m.TotalAmount = order.TotalAmount
	m.Status = order.Status
	m.Note = order.Note
	m.CreatedAt = order.CreatedAt
	m.UpdatedAt = order.UpdatedAt
}
