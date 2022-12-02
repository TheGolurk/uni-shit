package internal

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func (s *Service) CreateTypeUser(c echo.Context) error {

	return c.JSON(http.StatusOK, "created")
}

func (s *Service) DeleteTypeUser(c echo.Context) error {

	return c.JSON(http.StatusOK, "deleted")
}

func (s *Service) ModifyTypeUser(c echo.Context) error {

	return c.JSON(http.StatusOK, "updated")

}

func (s *Service) GetTypeUser(c echo.Context) error {
	return c.JSON(http.StatusOK, "deleted")
}
