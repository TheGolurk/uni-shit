package internal

import (
	"log"
	"net/http"
	"upemor.com/shit-project/models"

	"github.com/labstack/echo/v4"
)

func (s *Service) CreateTypeUser(c echo.Context) error {
	var (
		user models.TYpeUser
	)

	if err := c.Bind(&user); err != nil {
		log.Println(err)
		return c.JSON(http.StatusBadRequest, "")
	}

	_, err := s.db.Exec(TipoUser_Create, user.TipoUsuario)
	if err != nil {
		log.Println(err)
		return c.JSON(http.StatusBadRequest, "")
	}

	return c.JSON(http.StatusOK, "created")
}

func (s *Service) DeleteTypeUser(c echo.Context) error {
	id := c.QueryParam("id")
	if id == "" {
		return c.JSON(http.StatusBadRequest, "id faltante")
	}

	_, err := s.db.Exec(TipoUser_Delete, id)

	if err != nil {
		log.Println(err)
		return c.JSON(http.StatusBadRequest, "")
	}

	return c.JSON(http.StatusOK, "deleted")
}

func (s *Service) ModifyTypeUser(c echo.Context) error {
	var (
		user models.TYpeUser
	)

	if err := c.Bind(&user); err != nil {
		log.Println(err)
		return c.JSON(http.StatusBadRequest, "")
	}

	_, err := s.db.Exec(TipoUser_Update, user.TipoUsuario, user.ID)

	if err != nil {
		log.Println(err)
		return c.JSON(http.StatusBadRequest, "")
	}

	return c.JSON(http.StatusOK, "updated")
}

func (s *Service) GetTypeUser(c echo.Context) error {
	var (
		user  models.TYpeUser
		users []models.TYpeUser
	)

	rows, err := s.db.Query(TipoUser_get)
	if err != nil {
		log.Println(err)
		return c.JSON(http.StatusBadRequest, "")
	}

	defer rows.Close()

	for rows.Next() {
		if err = rows.Scan(
			&user.ID,
			&user.TipoUsuario,
		); err != nil {
			log.Println(err)
			return c.JSON(http.StatusInternalServerError, "")
		}
		users = append(users, user)
	}

	return c.JSON(http.StatusOK, users)
}
