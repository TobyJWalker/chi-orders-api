package model

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// create an order struct
type Order struct {
	gorm.Model

	CustomerID uuid.UUID `json:"customer_id"`
	LineItems []LineItem `json:"line_items" gorm:"many2many:order_items;"`
	CreatedAt *time.Time `json:"created_at"`
	ShippedAt *time.Time `json:"shipped_at"`
	CompletedAt *time.Time `json:"completed_at"`
}

// line item struct
type LineItem struct {
	gorm.Model
	ItemID uuid.UUID `json:"item_id"`
	Quantity uint `json:"quantity"`
	Price uint `json:"price"`
}

