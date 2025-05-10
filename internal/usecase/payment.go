package usecase

import (
	"context"
	"go-clean/internal/model"
	"go-clean/internal/repository"
	"go-clean/internal/repository/dto"
)

type payment struct {
	repoPayment repository.Payment
}

func NewPayment(repoPayment repository.Payment) *payment {
	return &payment{
		repoPayment: repoPayment,
	}
}
func (p *payment) Create(ctx context.Context, param *model.PaymentCreateRequest) (*model.Payment, error) {
	createdPayment, err := p.repoPayment.Create(ctx, &dto.PaymentCreateRequest{
		OrderID:     param.OrderID,
		Amount:      param.Amount,
		PaymentType: param.PaymentType,
	})
	if err != nil {
		return nil, err
	}
	return &model.Payment{
		ID:          createdPayment.ID,
		OrderID:     createdPayment.OrderID,
		Amount:      createdPayment.Amount,
		PaymentType: createdPayment.PaymentType,
		Status:      createdPayment.Status,
		CreatedAt:   createdPayment.CreatedAt,
		UpdatedAt:   createdPayment.UpdatedAt,
	}, err
}

func (p *payment) GetByID(ctx context.Context, id int64) (*model.Payment, error) {
	payment, err := p.repoPayment.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}
	return &model.Payment{
		ID:          payment.ID,
		OrderID:     payment.OrderID,
		Amount:      payment.Amount,
		PaymentType: payment.PaymentType,
		Status:      payment.Status,
		CreatedAt:   payment.CreatedAt,
		UpdatedAt:   payment.UpdatedAt,
	}, nil
}
func (p *payment) GetByOrderID(ctx context.Context, orderID int64) (*model.Payment, error) {
	payment, err := p.repoPayment.GetByOrderID(ctx, orderID)
	if err != nil {
		return nil, err
	}
	return &model.Payment{
		ID:          payment.ID,
		OrderID:     payment.OrderID,
		Amount:      payment.Amount,
		PaymentType: payment.PaymentType,
		Status:      payment.Status,
		CreatedAt:   payment.CreatedAt,
		UpdatedAt:   payment.UpdatedAt,
	}, nil
}
func (p *payment) Update(ctx context.Context, id int64, param *model.PaymentUpdate) (*model.Payment, error) {
	updatedPayment, err := p.repoPayment.Update(ctx, id, &dto.PaymentUpdate{
		Status: param.Status,
	})
	if err != nil {
		return nil, err
	}
	return &model.Payment{
		ID:          updatedPayment.ID,
		OrderID:     updatedPayment.OrderID,
		Amount:      updatedPayment.Amount,
		PaymentType: updatedPayment.PaymentType,
		Status:      updatedPayment.Status,
		CreatedAt:   updatedPayment.CreatedAt,
		UpdatedAt:   updatedPayment.UpdatedAt,
	}, nil
}
func (p *payment) List(ctx context.Context) ([]*model.Payment, error) {
	payments, err := p.repoPayment.List(ctx)
	if err != nil {
		return nil, err
	}
	var result []*model.Payment
	for _, payment := range payments {
		result = append(result, &model.Payment{
			ID:          payment.ID,
			OrderID:     payment.OrderID,
			Amount:      payment.Amount,
			PaymentType: payment.PaymentType,
			Status:      payment.Status,
			CreatedAt:   payment.CreatedAt,
			UpdatedAt:   payment.UpdatedAt,
		})
	}
	return result, nil
}
