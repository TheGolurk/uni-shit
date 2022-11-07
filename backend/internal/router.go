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

	// Login
	e.POST("/login", service.Login)

	// User Access management
	e.POST("/user/accesohora", service.CreateUserAccess)
}
