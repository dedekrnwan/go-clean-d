package repository

import (
	"context"
	"go-clean/internal/repository/dto"
)

// Payment defines the payment repository interface
type Payment interface {
	Create(ctx context.Context, payment *dto.PaymentCreateRequest) (*dto.Payment, error)
	GetByID(ctx context.Context, id int64) (*dto.Payment, error)
	GetByOrderID(ctx context.Context, orderID int64) (*dto.Payment, error)
	Update(ctx context.Context, id int64, payment *dto.PaymentUpdate) (*dto.Payment, error)
	List(ctx context.Context) ([]*dto.Payment, error)
}
