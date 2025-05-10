package model

// Payment represents payment data transfer object
type Payment struct {
	ID          int64   `json:"id"`
	OrderID     int64   `json:"order_id"`
	Amount      float64 `json:"amount"`
	Status      string  `json:"status"`
	PaymentType string  `json:"payment_type"`
	CreatedAt   string  `json:"created_at"`
	UpdatedAt   string  `json:"updated_at"`
}

// PaymentCreate represents payment creation data
type PaymentCreateRequest struct {
	OrderID     int64   `json:"order_id"`
	Amount      float64 `json:"amount"`
	PaymentType string  `json:"payment_type"`
}

// PaymentUpdate represents payment update data
type PaymentUpdate struct {
	Status string `json:"status"`
}
