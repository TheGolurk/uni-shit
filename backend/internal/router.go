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

	e.POST("/login", service.Login)
	e.POST("/create/user", service.CreateUser)
}
