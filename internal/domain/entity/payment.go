package entity

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"time"
)

type Payment struct {
	ID            uint
	OrderID       uint
	UserID        uint
	Amount        float64
	Currency      string
	PaymentMethod string
	Status        string
	Gateway       string
	GatewayRef    string
	GatewayData   JSONB
	ExpiresAt     *time.Time
	PaidAt        *time.Time
	CreatedAt     time.Time
	UpdatedAt     time.Time
}

// JSONB type for storing JSON data
type JSONB map[string]interface{}

func (j JSONB) Value() (driver.Value, error) {
	if j == nil {
		return nil, nil
	}
	return json.Marshal(j)
}

func (j *JSONB) Scan(value interface{}) error {
	if value == nil {
		*j = nil
		return nil
	}

	var bytes []byte
	switch v := value.(type) {
	case []byte:
		bytes = v
	case string:
		bytes = []byte(v)
	default:
		return errors.New("failed to unmarshal JSONB value")
	}

	return json.Unmarshal(bytes, j)
}

// Business methods
func (p *Payment) IsPending() bool {
	return p.Status == "pending"
}

func (p *Payment) IsSuccess() bool {
	return p.Status == "success"
}

func (p *Payment) IsFailed() bool {
	return p.Status == "failed"
}

func (p *Payment) IsExpired() bool {
	if p.ExpiresAt == nil {
		return false
	}
	return time.Now().After(*p.ExpiresAt)
}

func (p *Payment) CanBePaid() bool {
	return p.IsPending() && !p.IsExpired()
}

func (p *Payment) MarkAsSuccess() {
	now := time.Now()
	p.Status = "success"
	p.PaidAt = &now
}

func (p *Payment) MarkAsFailed() {
	p.Status = "failed"
}

func (p *Payment) MarkAsExpired() {
	p.Status = "expired"
}

func (p *Payment) SetGatewayData(data map[string]interface{}) {
	p.GatewayData = JSONB(data)
}
