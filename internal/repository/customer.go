package repository

import (
	"context"
	"go-clean/internal/repository/dto"
)

type Customer interface {
	// CreateCustomer creates a new customer in the database
	Create(ctx context.Context, customer *dto.CustomerCreateRequest) (*dto.Customer, error)
	// GetCustomerByID retrieves a customer by ID from the database
	GetByID(ctx context.Context, customerID string) (*dto.Customer, error)
	// UpdateCustomer updates an existing customer in the database
	Update(ctx context.Context, customer *dto.CustomerUpdateRequest) (*dto.Customer, error)
	// DeleteCustomer deletes a customer from the database
	Delete(ctx context.Context, customerID string) error
	// GetCustomerByEmail retrieves a customer by email from the database
	GetByEmail(ctx context.Context, email string) (*dto.Customer, error)
	// GetCustomerByPhone retrieves a customer by phone number from the database
	GetByPhone(ctx context.Context, hone string) (*dto.Customer, error)
	// ListCustomers retrieves a list of customers from the database
	List(ctx context.Context) ([]*dto.Customer, error)
}
