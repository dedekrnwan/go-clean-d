package pg

import (
	"context"
	"go-clean/internal/repository/dto"
)

type customer struct {
	connection any
}

func NewCustomer(connection any) *customer {
	return &customer{
		connection: connection,
	}
}
func (c *customer) Create(ctx context.Context, customer *dto.CustomerCreateRequest) (*dto.Customer, error) {
	// Implement the logic to create a customer in the PostgreSQL database
	// using the provided connection and return the created customer entity.
	return nil, nil
}
func (c *customer) GetByID(ctx context.Context, id string) (*dto.Customer, error) {
	// Implement the logic to retrieve a customer by ID from the PostgreSQL database
	// using the provided connection and return the customer entity.
	return nil, nil
}
func (c *customer) Update(ctx context.Context, customer *dto.CustomerUpdateRequest) (*dto.Customer, error) {
	// Implement the logic to update a customer in the PostgreSQL database
	// using the provided connection and return the updated customer entity.
	return nil, nil
}
func (c *customer) Delete(ctx context.Context, id string) error {
	// Implement the logic to delete a customer by ID from the PostgreSQL database
	// using the provided connection.
	return nil
}
func (c *customer) List(ctx context.Context) ([]*dto.Customer, error) {
	// Implement the logic to retrieve a list of customers from the PostgreSQL database
	// using the provided connection and return the list of customer entities.
	return nil, nil
}
func (c *customer) GetByEmail(ctx context.Context, email string) (*dto.Customer, error) {
	// Implement the logic to retrieve a customer by email from the PostgreSQL database
	// using the provided connection and return the customer entity.
	return nil, nil
}
func (c *customer) GetByPhone(ctx context.Context, phone string) (*dto.Customer, error) {
	// Implement the logic to retrieve a customer by phone from the PostgreSQL database
	// using the provided connection and return the customer entity.
	return nil, nil
}
