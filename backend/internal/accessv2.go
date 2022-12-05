package internal

import (
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
	"upemor.com/shit-project/models"
)

func (s *Service) CreateAccessV2(c echo.Context) error {
	var (
		access models.Acceso
	)

	if err := c.Bind(&access); err != nil {
		log.Println(err)
		return c.JSON(http.StatusBadRequest, "")
	}

	_, err := s.db.Exec(Acceso_create, access.IdTipo, access.Tablas)
	if err != nil {
		log.Println(err)
		return c.JSON(http.StatusBadRequest, "")
	}

	return c.JSON(http.StatusOK, "created")
}

func (s *Service) DeleteAccessV2(c echo.Context) error {
	id := c.QueryParam("id")
	if id == "" {
		return c.JSON(http.StatusBadRequest, "")
	}

	_, err := s.db.Exec(Acceso_delete, id)

	if err != nil {
		log.Println(err)
		return c.JSON(http.StatusBadRequest, "")
	}

	return c.JSON(http.StatusOK, "created")

}

func (s *Service) UpdateAccessV2(c echo.Context) error {
	var (
		access models.Acceso
	)

	if err := c.Bind(&access); err != nil {
		log.Println(err)
		return c.JSON(http.StatusBadRequest, "")
	}

	_, err := s.db.Exec(Acceso_update, access.IdTipo, access.Tablas, access.ID)
	if err != nil {
		log.Println(err)
		return c.JSON(http.StatusBadRequest, "")
	}

	return c.JSON(http.StatusOK, "updated")
}

func (s *Service) GetAccessV2(c echo.Context) error {
	var (
		access    models.Acceso
		accessarr []models.Acceso
	)

	rows, err := s.db.Query(Acceso_get)
	if err != nil {
		log.Println(err)
		return c.JSON(http.StatusBadRequest, "")
	}

	defer rows.Close()

	for rows.Next() {
		if err = rows.Scan(
			&access.ID,
			&access.IdTipo,
			&access.Tablas,
		); err != nil {
			log.Println(err)
			return c.JSON(http.StatusInternalServerError, "")
		}
		accessarr = append(accessarr, access)
	}

	return c.JSON(http.StatusOK, accessarr)
}
