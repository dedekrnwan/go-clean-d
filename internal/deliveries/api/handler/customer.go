package handler

import (
	"go-clean/internal/model"
	"go-clean/internal/repository/dto"
	"go-clean/internal/usecase"
	"go-clean/pkg/api"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/super-saga/go-x/clue"
)

type (
	customer struct {
		usecase *usecase.Customer
	}
	Customer interface {
		Create(customer *dto.CustomerCreateRequest) error
		GetByID(customerID string) (*dto.Customer, error)
		Update(customer *dto.CustomerUpdateRequest) error
		Delete(customerID string) error
		GetByEmail(email string) (*dto.Customer, error)
		GetByPhone(phone string) (*dto.Customer, error)
		List() ([]*dto.Customer, error)
	}
)

func NewCustomer(usecase *usecase.Customer) *customer {
	return &customer{
		usecase: usecase,
	}
}

func (h *customer) Create(c echo.Context) error {
	var (
		ctx     = c.Request().Context()
		payload = &dto.CustomerCreateRequest{}
	)
	err := api.BindRequest(c, payload, false)
	if err != nil {
		return clue.Build(http.StatusBadRequest, "00", nil, "Bad Request").Send(c)
	}

	out, err := h.usecase.Create(ctx, &model.CustomerCreateRequest{
		Name:        payload.Name,
		Email:       payload.Email,
		Phone:       payload.Phone,
		Address:     payload.Address,
		City:        payload.City,
		State:       payload.State,
		ZipCode:     payload.ZipCode,
		Country:     payload.Country,
		DateOfBirth: payload.DateOfBirth,
	})
	if err != nil {
		return clue.CoverBuilder(err, nil).Send(c)
	}

	return clue.Build(http.StatusOK, "00", &dto.Customer{
		ID:          out.ID,
		Name:        out.Name,
		Email:       out.Email,
		Phone:       out.Phone,
		Address:     out.Address,
		City:        out.City,
		State:       out.State,
		ZipCode:     out.ZipCode,
		Country:     out.Country,
		DateOfBirth: out.DateOfBirth,
	}, "Success").Send(c)

}
func (h *customer) GetByID(c echo.Context) error {
	var (
		ctx = c.Request().Context()
	)

	out, err := h.usecase.GetByID(ctx, "1")
	if err != nil {
		return clue.CoverBuilder(err, nil).Send(c)
	}

	return clue.Build(http.StatusOK, "00", &dto.Customer{
		ID:          out.ID,
		Name:        out.Name,
		Email:       out.Email,
		Phone:       out.Phone,
		Address:     out.Address,
		City:        out.City,
		State:       out.State,
		ZipCode:     out.ZipCode,
		Country:     out.Country,
		DateOfBirth: out.DateOfBirth,
	}, "Success").Send(c)
}

func (h *customer) Update(c echo.Context) error {
	var (
		ctx     = c.Request().Context()
		payload = &dto.CustomerUpdateRequest{}
	)
	err := api.BindRequest(c, payload, false)
	if err != nil {
		return clue.Build(http.StatusBadRequest, "00", nil, "Bad Request").Send(c)
	}

	out, err := h.usecase.Update(ctx, &model.CustomerUpdateRequest{
		ID:          payload.ID,
		Name:        payload.Name,
		Email:       payload.Email,
		Phone:       payload.Phone,
		Address:     payload.Address,
		City:        payload.City,
		State:       payload.State,
		ZipCode:     payload.ZipCode,
		Country:     payload.Country,
		DateOfBirth: payload.DateOfBirth,
	})
	if err != nil {
		return clue.CoverBuilder(err, nil).Send(c)
	}

	return clue.Build(http.StatusOK, "00", &dto.Customer{
		ID:          out.ID,
		Name:        out.Name,
		Email:       out.Email,
		Phone:       out.Phone,
		Address:     out.Address,
		City:        out.City,
		State:       out.State,
		ZipCode:     out.ZipCode,
		Country:     out.Country,
		DateOfBirth: out.DateOfBirth,
	}, "Success").Send(c)
}

func (h *customer) Delete(c echo.Context) error {
	var (
		ctx = c.Request().Context()
	)

	err := h.usecase.Delete(ctx, "1")
	if err != nil {
		return clue.CoverBuilder(err, nil).Send(c)
	}

	return clue.Build(http.StatusOK, "00", nil, "Success").Send(c)
}

func (h *customer) GetByEmail(c echo.Context) error {
	var (
		ctx = c.Request().Context()
	)

	out, err := h.usecase.GetByEmail(ctx, "1")
	if err != nil {
		return clue.CoverBuilder(err, nil).Send(c)
	}

	return clue.Build(http.StatusOK, "00", &dto.Customer{
		ID:          out.ID,
		Name:        out.Name,
		Email:       out.Email,
		Phone:       out.Phone,
		Address:     out.Address,
		City:        out.City,
		State:       out.State,
		ZipCode:     out.ZipCode,
		Country:     out.Country,
		DateOfBirth: out.DateOfBirth,
	}, "Success").Send(c)
}

func (h *customer) GetByPhone(c echo.Context) error {
	var (
		ctx = c.Request().Context()
	)

	out, err := h.usecase.GetByPhone(ctx, "1")
	if err != nil {
		return clue.CoverBuilder(err, nil).Send(c)
	}

	return clue.Build(http.StatusOK, "00", &dto.Customer{
		ID:          out.ID,
		Name:        out.Name,
		Email:       out.Email,
		Phone:       out.Phone,
		Address:     out.Address,
		City:        out.City,
		State:       out.State,
		ZipCode:     out.ZipCode,
		Country:     out.Country,
		DateOfBirth: out.DateOfBirth,
	}, "Success").Send(c)
}
func (h *customer) List(c echo.Context) error {
	var (
		ctx = c.Request().Context()
	)
	customers, err := h.usecase.List(ctx)
	if err != nil {
		return clue.CoverBuilder(err, nil).Send(c)
	}
	var items []*dto.Customer
	for _, item := range customers {
		items = append(items, &dto.Customer{
			ID:          item.ID,
			Name:        item.Name,
			Email:       item.Email,
			Phone:       item.Phone,
			Address:     item.Address,
			City:        item.City,
			State:       item.State,
			ZipCode:     item.ZipCode,
			Country:     item.Country,
			DateOfBirth: item.DateOfBirth,
		})
	}
	return clue.Build(http.StatusOK, "00", items, "Success").Send(c)
}
