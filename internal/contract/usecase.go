package contract

import "go-clean/internal/usecase"

type Usecase struct {
	Customer *usecase.Customer
}

func NewUsecase(repo *Repository) *Usecase {
	return &Usecase{
		Customer: usecase.NewCustomer(repo.Customer),
	}
}
