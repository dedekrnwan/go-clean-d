package api

import (
	"fmt"

	"github.com/labstack/echo/v4"
	xlog "github.com/super-saga/go-x/log"
)

func BindRequest(c echo.Context, payload interface{}, skipValidation bool) error {
	ctx := c.Request().Context()

	b := &echo.DefaultBinder{}
	if err := b.BindHeaders(c, payload); err != nil {
		xlog.Error(ctx, err.Error())
		return fmt.Errorf("%s: %s", "Invalid Format", err.Error())
	}
	if err := c.Bind(payload); err != nil {
		xlog.Error(ctx, err.Error())
		return fmt.Errorf("%s: %s", "Invalid Format", err.Error())
	}

	if !skipValidation {
		if err := c.Validate(payload); err != nil {
			xlog.Error(ctx, err.Error())
			return fmt.Errorf("%s: %s", "Mandatory Field", err.Error())
		}
	}

	return nil
}
