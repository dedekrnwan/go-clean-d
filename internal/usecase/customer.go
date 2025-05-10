package usecase

import (
	"context"
	"go-clean/internal/model"
	"go-clean/internal/repository"
	"go-clean/internal/repository/dto"
)

type Customer struct {
	repoCustomer repository.Customer
}

func NewCustomer(repoCustomer repository.Customer) *Customer {
	return &Customer{
		repoCustomer: repoCustomer,
	}
}

func (c *Customer) Create(ctx context.Context, customer *model.CustomerCreateRequest) (*model.Customer, error) {
	customerEntity, err := c.repoCustomer.Create(ctx, &dto.CustomerCreateRequest{
		Name:        customer.Name,
		Email:       customer.Email,
		Phone:       customer.Phone,
		Address:     customer.Address,
		City:        customer.City,
		State:       customer.State,
		ZipCode:     customer.ZipCode,
		Country:     customer.Country,
		DateOfBirth: customer.DateOfBirth,
	})
	if err != nil {
		return nil, err
	}
	return &model.Customer{
		ID:          customerEntity.ID,
		Name:        customerEntity.Name,
		Email:       customerEntity.Email,
		Phone:       customerEntity.Phone,
		Address:     customerEntity.Address,
		City:        customerEntity.City,
		State:       customerEntity.State,
		ZipCode:     customerEntity.ZipCode,
		Country:     customerEntity.Country,
		DateOfBirth: customerEntity.DateOfBirth,
	}, nil
}
func (c *Customer) GetByID(ctx context.Context, customerID string) (*model.Customer, error) {
	customer, err := c.repoCustomer.GetByID(ctx, customerID)
	if err != nil {
		return nil, err
	}
	return &model.Customer{
		ID:          customer.ID,
		Name:        customer.Name,
		Email:       customer.Email,
		Phone:       customer.Phone,
		Address:     customer.Address,
		City:        customer.City,
		State:       customer.State,
		ZipCode:     customer.ZipCode,
		Country:     customer.Country,
		DateOfBirth: customer.DateOfBirth,
	}, nil
}
func (c *Customer) Update(ctx context.Context, customer *model.CustomerUpdateRequest) (*model.Customer, error) {
	customerEntity, err := c.repoCustomer.Update(ctx, &dto.CustomerUpdateRequest{
		ID:          customer.ID,
		Name:        customer.Name,
		Email:       customer.Email,
		Phone:       customer.Phone,
		Address:     customer.Address,
		City:        customer.City,
		State:       customer.State,
		ZipCode:     customer.ZipCode,
		Country:     customer.Country,
		DateOfBirth: customer.DateOfBirth,
	})
	if err != nil {
		return nil, err
	}
	return &model.Customer{
		ID:          customerEntity.ID,
		Name:        customerEntity.Name,
		Email:       customerEntity.Email,
		Phone:       customerEntity.Phone,
		Address:     customerEntity.Address,
		City:        customerEntity.City,
		State:       customerEntity.State,
		ZipCode:     customerEntity.ZipCode,
		Country:     customer.Country,
		DateOfBirth: customer.DateOfBirth,
	}, nil
}
func (c *Customer) Delete(ctx context.Context, customerID string) error {
	return c.repoCustomer.Delete(ctx, customerID)
}
func (c *Customer) List(ctx context.Context) ([]*model.Customer, error) {
	customers, err := c.repoCustomer.List(ctx)
	if err != nil {
		return nil, err
	}
	var result []*model.Customer
	for _, customer := range customers {
		result = append(result, &model.Customer{
			ID:          customer.ID,
			Name:        customer.Name,
			Email:       customer.Email,
			Phone:       customer.Phone,
			Address:     customer.Address,
			City:        customer.City,
			State:       customer.State,
			ZipCode:     customer.ZipCode,
			Country:     customer.Country,
			DateOfBirth: customer.DateOfBirth,
		})
	}
	return result, nil
}
func (c *Customer) GetByEmail(ctx context.Context, email string) (*model.Customer, error) {
	customer, err := c.repoCustomer.GetByEmail(ctx, email)
	if err != nil {
		return nil, err
	}
	return &model.Customer{
		ID:          customer.ID,
		Name:        customer.Name,
		Email:       customer.Email,
		Phone:       customer.Phone,
		Address:     customer.Address,
		City:        customer.City,
		State:       customer.State,
		ZipCode:     customer.ZipCode,
		Country:     customer.Country,
		DateOfBirth: customer.DateOfBirth,
	}, nil
}
func (c *Customer) GetByPhone(ctx context.Context, phone string) (*model.Customer, error) {
	customer, err := c.repoCustomer.GetByPhone(ctx, phone)
	if err != nil {
		return nil, err
	}
	return &model.Customer{
		ID:          customer.ID,
		Name:        customer.Name,
		Email:       customer.Email,
		Phone:       customer.Phone,
		Address:     customer.Address,
		City:        customer.City,
		State:       customer.State,
		ZipCode:     customer.ZipCode,
		Country:     customer.Country,
		DateOfBirth: customer.DateOfBirth,
	}, nil
}
