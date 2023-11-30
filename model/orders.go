package model

import (
	"time"

	"github.com/google/uuid"
)

// create an order struct
type Order struct {
	OrderID uint64 `json:"order_id" gorm:"primaryKey"`
	CustomerID uuid.UUID `json:"customer_id"`
	LineItems []LineItem `json:"line_items" gorm:"foreignKey:OrderID"`
	CreatedAt *time.Time `json:"created_at"`
	ShippedAt *time.Time `json:"shipped_at"`
	CompletedAt *time.Time `json:"completed_at"`
}

// line item struct
type LineItem struct {
	ItemID uuid.UUID
	Quantity uint
	Price uint
}

