package order

import (
	"chi-orders-api/model"
	"context"
	"fmt"

	"gorm.io/gorm"
)

// repository struct
type PostgresRepo struct {
	Client *gorm.DB
}

// insert method
func (r *PostgresRepo) Insert(ctx context.Context, order model.Order) error {

	// create order
	err := r.Client.Create(&order).Error
	if err != nil {
		return fmt.Errorf("failed to set: %w", err)
	}

	return nil
}

// find by id
func (r *PostgresRepo) FindByID(ctx context.Context, id uint64) (model.Order, error) {

	// init order object
	var order model.Order

	// find order
	err := r.Client.First(&order, id).Error
	if err != nil {
		return order, fmt.Errorf("failed to find: %w", err)
	}

	return order, nil

}

// delete by id
func (r *PostgresRepo) DeleteByID(ctx context.Context, id uint64) error {

	// delete order
	err := r.Client.Delete(&model.Order{}, id).Error
	if err != nil {
		return fmt.Errorf("failed to delete: %w", err)
	}

	return nil

}

// update
func (r *PostgresRepo) Update(ctx context.Context, order model.Order) error {

	// update order
	err := r.Client.Save(&order).Error
	if err != nil {
		return fmt.Errorf("failed to update: %w", err)
	}

	return nil

}

// find all
func (r *PostgresRepo) FindAll(ctx context.Context) ([]model.Order, error) {

	// init orders slice
	var orders []model.Order

	// find orders
	err := r.Client.Preload("LineItems").Find(&orders).Error
	if err != nil {
		return orders, fmt.Errorf("failed to find all: %w", err)
	}

	return orders, nil

}