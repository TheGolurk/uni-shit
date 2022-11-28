package internal

import (
	"database/sql"

	"github.com/labstack/echo/v4"
)

type Service struct {
	db *sql.DB
}

func SetRoutes(e *echo.Echo, db *sql.DB) {
	service := Service{db: db}

	// User
	e.POST("/create/user", service.CreateUser)
	e.DELETE("/user/delete", service.DeleteUser)
	e.PUT("/user/modify", service.ModifyUser)
	e.GET("/user/get", service.GetUser)

	// Login
	e.POST("/login", service.Login)

	// User Access management
	e.POST("/accesohora/create", service.CreateAccess)
	e.DELETE("/accesohora/delete", service.DeleteAccess)
	e.PUT("/accesohora/modify", service.UpdateAccess)
	e.GET("/accesohora/get", service.GetAccess)

	// Venta
	e.POST("/venta/create", service.CreateSell)
	e.DELETE("/venta/delete", service.DeleteSell)
	e.PUT("/venta/modify", service.UpdateSell)
	e.GET("/venta/get", service.GetSell)

	// PRODUCTO
	e.POST("/producto/create", service.CreateProduct)
	e.DELETE("/producto/delete", service.DeleteProduct)
	e.PUT("/producto/modify", service.UpdateProduct)
	e.GET("/producto/get", service.GetProduct)

	// CLIENTE
	e.POST("/cliente/create", service.CreateClient)
	e.DELETE("/cliente/delete", service.DeleteClient)
	e.PUT("/cliente/modify", service.UpdateClient)
	e.GET("/cliente/get", service.GetClient)

}
