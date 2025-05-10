package model

type Customer struct {
	// ID of the customer
	ID string `json:"id"`
	// Name of the customer
	Name string `json:"name"`
	// Email of the customer
	Email string `json:"email"`
	// Phone of the customer
	Phone string `json:"phone"`
	// Address of the customer
	Address string `json:"address"`
	// City of the customer
	City string `json:"city"`
	// State of the customer
	State string `json:"state"`
	// ZipCode of the customer
	ZipCode string `json:"zip_code"`
	// Country of the customer
	Country string `json:"country"`
	// DateOfBirth of the customer
	DateOfBirth string `json:"date_of_birth"`
}

type CustomerCreateRequest struct {
	// Name of the customer
	Name string `json:"name" validate:"required"`
	// Email of the customer
	Email string `json:"email" validate:"required,email"`
	// Phone of the customer
	Phone string `json:"phone" validate:"required"`
	// Address of the customer
	Address string `json:"address" validate:"required"`
	// City of the customer
	City string `json:"city" validate:"required"`
	// State of the customer
	State string `json:"state" validate:"required"`
	// ZipCode of the customer
	ZipCode string `json:"zip_code" validate:"required"`
	// Country of the customer
	Country string `json:"country" validate:"required"`
	// DateOfBirth of the customer
	DateOfBirth string `json:"date_of_birth" validate:"required"`
}

type CustomerUpdateRequest struct {
	// ID of the customer
	ID string `json:"id" validate:"required"`
	// Name of the customer
	Name string `json:"name" validate:"required"`
	// Email of the customer
	Email string `json:"email" validate:"required,email"`
	// Phone of the customer
	Phone string `json:"phone" validate:"required"`
	// Address of the customer
	Address string `json:"address" validate:"required"`
	// City of the customer
	City string `json:"city" validate:"required"`
	// State of the customer
	State string `json:"state" validate:"required"`
	// ZipCode of the customer
	ZipCode string `json:"zip_code" validate:"required"`
	// Country of the customer
	Country string `json:"country" validate:"required"`
	// DateOfBirth of the customer
	DateOfBirth string `json:"date_of_birth" validate:"required"`
}
