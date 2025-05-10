package midtrans

import (
	"context"
	"go-clean/internal/repository/dto"
)

type payment struct {
	connection any
}

func NewPayment(connection any) *payment {
	return &payment{
		connection: connection,
	}
}
func (p *payment) Create(ctx context.Context, payment *dto.PaymentCreateRequest) (*dto.Payment, error) {
	// Implement the logic to create a payment in the Xendit API
	// using the provided connection and return the created payment entity.
	return nil, nil
}
func (p *payment) GetByID(ctx context.Context, id int64) (*dto.Payment, error) {
	// Implement the logic to retrieve a payment by ID from the Xendit API
	// using the provided connection and return the payment entity.
	return nil, nil
}
func (p *payment) GetByOrderID(ctx context.Context, orderID int64) (*dto.Payment, error) {
	// Implement the logic to retrieve a payment by order ID from the Xendit API
	// using the provided connection and return the payment entity.
	return nil, nil
}
func (p *payment) Update(ctx context.Context, id int64, payment *dto.PaymentUpdate) (*dto.Payment, error) {
	// Implement the logic to update a payment in the Xendit API
	// using the provided connection and return the updated payment entity.
	return nil, nil
}
func (p *payment) List(ctx context.Context) ([]*dto.Payment, error) {
	// Implement the logic to retrieve a list of payments from the Xendit API
	// using the provided connection and return the list of payment entities.
	return nil, nil
}
