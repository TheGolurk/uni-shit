package internal

import (
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
	"upemor.com/shit-project/models"
)

func (s *Service) CreateAccess(c echo.Context) error {
	var (
		user models.User
	)

	if err := c.Bind(&user); err != nil {
		log.Println(err)
		return c.JSON(http.StatusInternalServerError, "")
	}

	pwd, err := hashAndSalt(user.Password)
	if err != nil {
		log.Println(err)
		return c.JSON(http.StatusInternalServerError, "")
	}

	_, err = s.db.Exec(`
	INSERT INTO USUARIO( USERNAME, PASS, NOMBRE, APELLIDO, TIPO_ID)
    	VALUES (?, ?, ?, ?, ?);
		`,
		user.Username,
		pwd,
		user.Nombre,
		user.Apellido,
		user.IdTipo,
	)
	if err != nil {
		log.Println(err)
		return c.JSON(http.StatusInternalServerError, "")
	}

	return c.JSON(http.StatusOK, "created")

}

func (s *Service) DeleteAccess(c echo.Context) error {

	return nil
}

func (s *Service) UpdateAccess(c echo.Context) error {

	return nil
}

func (s *Service) GetAccess(c echo.Context) error {

	return nil
}
