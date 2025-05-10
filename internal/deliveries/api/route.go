package api

import (
	"go-clean/internal/contract"
	"go-clean/internal/deliveries/api/handler"

	"github.com/labstack/echo/v4"
)

func Bind(g *echo.Group, dep *contract.Dependency) {
	customerHandler := handler.NewCustomer(dep.Usecase.Customer)

	//customer
	customer := g.Group("/customer")
	customer.POST("/create", customerHandler.Create)
	customer.GET("/get/:id", customerHandler.GetByID)
	customer.GET("/get", customerHandler.List)
	customer.PUT("/update/:id", customerHandler.Update)
	customer.DELETE("/delete/:id", customerHandler.Delete)
	customer.GET("/get-by-email", customerHandler.GetByEmail)
	customer.GET("/get-by-phone", customerHandler.GetByPhone)
}
