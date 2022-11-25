package internal

import (
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
	"upemor.com/shit-project/models"
)

func (s *Service) CreateAccess(c echo.Context) error {
	var (
		access models.UserAccess
	)

	if err := c.Bind(&access); err != nil {
		log.Println(err)
		return c.JSON(http.StatusInternalServerError, "")
	}

	_, err := s.db.Exec(Tipoacceso_create,
		access.IDTipo,
		access.Tablas,
		access.HoraInicio,
		access.HoraFinal,
	)

	if err != nil {
		log.Println(err)
		return c.JSON(http.StatusInternalServerError, "")
	}

	return c.JSON(http.StatusOK, "created")
}

func (s *Service) DeleteAccess(c echo.Context) error {
	var (
		access models.UserAccess
	)

	if err := c.Bind(&access); err != nil {
		log.Println(err)
		return c.JSON(http.StatusInternalServerError, "")
	}

	_, err := s.db.Exec(Tipoacceso_delete, access.IDTipo)

	if err != nil {
		log.Println(err)
		return c.JSON(http.StatusInternalServerError, "")
	}

	return c.JSON(http.StatusOK, "created")

}

func (s *Service) UpdateAccess(c echo.Context) error {
	var (
		user models.User
	)

	if err := c.Bind(&user); err != nil {
		log.Println(err)
		return c.JSON(http.StatusInternalServerError, "")
	}

	_, err := s.db.Exec("")

	if err != nil {
		log.Println(err)
		return c.JSON(http.StatusInternalServerError, "")
	}

	return c.JSON(http.StatusOK, "created")

}

func (s *Service) GetAccess(c echo.Context) error {

	return nil
}
