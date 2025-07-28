package entity

import (
	"time"
)

type Order struct {
	ID          uint
	UserID      uint
	OrderNumber string
	TotalAmount float64
	Status      string
	Note        string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

// Business methods
func (o *Order) IsPending() bool {
	return o.Status == "pending"
}

func (o *Order) IsPaid() bool {
	return o.Status == "paid"
}

func (o *Order) IsCancelled() bool {
	return o.Status == "cancelled"
}

func (o *Order) IsCompleted() bool {
	return o.Status == "completed"
}

func (o *Order) CanBeCancelled() bool {
	return o.IsPending()
}

func (o *Order) CanBeCompleted() bool {
	return o.IsPaid()
}

func (o *Order) MarkAsPaid() {
	o.Status = "paid"
}

func (o *Order) MarkAsCancelled() {
	o.Status = "cancelled"
}

func (o *Order) MarkAsCompleted() {
	o.Status = "completed"
}

func (o *Order) GetStatusDisplay() string {
	switch o.Status {
	case "pending":
		return "Pending"
	case "paid":
		return "Paid"
	case "cancelled":
		return "Cancelled"
	case "completed":
		return "Completed"
	default:
		return "Unknown"
	}
}
