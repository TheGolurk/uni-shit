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
		return c.JSON(http.StatusBadRequest, "")
	}

	_, err := s.db.Exec(Tipoacceso_create,
		access.IDTipo,
		access.HoraInicio,
		access.HoraFinal,
	)

	if err != nil {
		log.Println(err)
		return c.JSON(http.StatusBadRequest, "")
	}

	return c.JSON(http.StatusOK, "created")
}

func (s *Service) DeleteAccess(c echo.Context) error {
	id := c.QueryParam("id")
	if id == "" {
		return c.JSON(http.StatusBadRequest, "id faltante")
	}

	_, err := s.db.Exec(Tipoacceso_delete, id)

	if err != nil {
		log.Println(err)
		return c.JSON(http.StatusBadRequest, "")
	}

	return c.JSON(http.StatusOK, "created")

}

func (s *Service) UpdateAccess(c echo.Context) error {
	var (
		access models.UserAccess
	)

	if err := c.Bind(&access); err != nil {
		log.Println(err)
		return c.JSON(http.StatusBadRequest, "")
	}

	_, err := s.db.Exec(Tipoacceso_update,
		access.HoraInicio,
		access.HoraFinal,
		access.IDTipo,
		access.ID,
	)

	if err != nil {
		log.Println(err)
		return c.JSON(http.StatusBadRequest, "")
	}

	return c.JSON(http.StatusOK, "updated")
}

func (s *Service) GetAccess(c echo.Context) error {
	var (
		access    models.UserAccess
		accessarr []models.UserAccess
	)

	rows, err := s.db.Query(Tipoacceso_select)
	if err != nil {
		log.Println(err)
		return c.JSON(http.StatusBadRequest, "")
	}

	defer rows.Close()

	for rows.Next() {
		if err = rows.Scan(
			&access.ID,
			&access.IDTipo,
			&access.HoraInicio,
			&access.HoraFinal,
		); err != nil {
			log.Println(err)
			return c.JSON(http.StatusInternalServerError, "")
		}
		accessarr = append(accessarr, access)
	}

	return c.JSON(http.StatusOK, accessarr)
}
