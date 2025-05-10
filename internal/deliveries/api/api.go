package api

import (
	"context"
	"go-clean/internal/contract"
	"net/http"

	"github.com/labstack/echo"
	"github.com/super-saga/go-x/clue"
	"github.com/super-saga/go-x/graceful"
)

type (
	Http interface {
		Prepare() (graceful.ProcessStarter, graceful.ProcessStopper)
	}
	api struct {
		e     *echo.Echo
		ctrct *contract.Dependency
	}
)

func New(c *contract.Dependency) Http {
	e := echo.New()
	e.HideBanner = true
	return &api{
		e:     e,
		ctrct: c,
	}
}

func (a *api) NewRoute(c *contract.Dependency) {
	a.e.GET("/", func(c echo.Context) error {
		message := "Welcome to go-clean version 1.0"
		return c.JSON(http.StatusOK, clue.Build(http.StatusOK, "00", nil, message))
	})
}

func (a *api) Prepare() (graceful.ProcessStarter, graceful.ProcessStopper) {
	a.NewRoute(a.ctrct)
	return func() error {
			return a.e.Start(":3000")
		}, func(ctx context.Context) error {
			return a.e.Shutdown(ctx)
		}
}
