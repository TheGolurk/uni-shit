package internal

import (
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
	"upemor.com/shit-project/models"
)

func (s *Service) CreateClient(c echo.Context) error {
	var (
		client models.Cliente
	)

	if err := c.Bind(&client); err != nil {
		log.Println(err)
		return c.JSON(http.StatusInternalServerError, "")
	}

	_, err := s.db.Exec(
		Cliente_create,
		client.Nombre,
		client.Apellido,
		client.Direccion,
		client.Estado,
	)

	if err != nil {
		log.Println(err)
		return c.JSON(http.StatusInternalServerError, "")
	}

	return c.JSON(http.StatusOK, "created")
}

func (s *Service) DeleteClient(c echo.Context) error {
	var (
		client models.Cliente
	)

	if err := c.Bind(&client); err != nil {
		log.Println(err)
		return c.JSON(http.StatusInternalServerError, "")
	}

	_, err := s.db.Exec(Cliente_delete, client.Nombre)

	if err != nil {
		log.Println(err)
		return c.JSON(http.StatusInternalServerError, "")
	}

	return c.JSON(http.StatusOK, "deleted")
}

func (s *Service) UpdateClient(c echo.Context) error {
	var (
		client models.Cliente
	)

	if err := c.Bind(&client); err != nil {
		log.Println(err)
		return c.JSON(http.StatusInternalServerError, "")
	}

	_, err := s.db.Exec(Cliente_update,
		client.Nombre,
		client.Apellido,
		client.Direccion,
		client.Estado,
		client.IdCliente,
	)

	if err != nil {
		log.Println(err)
		return c.JSON(http.StatusInternalServerError, "")
	}

	return c.JSON(http.StatusOK, "updated")

}

func (s *Service) GetClient(c echo.Context) error {

	return nil
}
