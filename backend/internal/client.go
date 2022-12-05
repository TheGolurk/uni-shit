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
		return c.JSON(http.StatusBadRequest, "")
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
		return c.JSON(http.StatusBadRequest, "")
	}

	return c.JSON(http.StatusOK, "created")
}

func (s *Service) DeleteClient(c echo.Context) error {
	id := c.QueryParam("id")
	if id == "" {
		log.Println("id no viene")
		return c.JSON(http.StatusBadRequest, "")
	}

	_, err := s.db.Exec(Cliente_delete, id)

	if err != nil {
		log.Println(err)
		return c.JSON(http.StatusBadRequest, "")
	}

	return c.JSON(http.StatusOK, "deleted")
}

func (s *Service) UpdateClient(c echo.Context) error {
	var (
		client models.Cliente
	)

	if err := c.Bind(&client); err != nil {
		log.Println(err)
		return c.JSON(http.StatusBadRequest, "")
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
		return c.JSON(http.StatusBadRequest, "")
	}

	return c.JSON(http.StatusOK, "updated")

}

func (s *Service) GetClient(c echo.Context) error {
	var (
		client    models.Cliente
		clientarr []models.Cliente
	)

	rows, err := s.db.Query(Cliente_select)
	if err != nil {
		log.Println(err)
		return c.JSON(http.StatusInternalServerError, "")
	}

	defer rows.Close()

	for rows.Next() {
		if err = rows.Scan(
			&client.IdCliente,
			&client.Nombre,
			&client.Apellido,
			&client.Direccion,
			&client.Estado,
		); err != nil {
			log.Println(err)
			return c.JSON(http.StatusInternalServerError, "")
		}
		clientarr = append(clientarr, client)
	}

	return c.JSON(http.StatusOK, clientarr)
}
