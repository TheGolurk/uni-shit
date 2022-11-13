package main

import (
	"github.com/labstack/echo/v4/middleware"
	"upemor.com/shit-project/config"
	"upemor.com/shit-project/internal"

	"github.com/labstack/echo/v4"
)

func main() {
	db := config.GetDB()

	e := echo.New()

	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "method=${method}, uri=${uri}, status=${status}\n",
	}))

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{},
	}))

	internal.SetRoutes(e, db)

	e.Logger.Fatal(e.Start(":8070"))
}
