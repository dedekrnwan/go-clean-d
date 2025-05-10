package contract

import (
	"go-clean/internal/repository"
	"go-clean/internal/repository/midtrans"
	"go-clean/internal/repository/pg"
)

type Repository struct {
	Customer repository.Customer
	Payment  repository.Payment
}

func NewRepository() *Repository {
	return &Repository{
		Customer: pg.NewCustomer("connection ceritanya"),
		Payment:  midtrans.NewPayment("baseurl ceritanya"),
	}
}
