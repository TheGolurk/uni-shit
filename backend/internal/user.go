package internal

import (
	"errors"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
	"upemor.com/shit-project/models"
)

func (s *Service) Login(c echo.Context) error {
	var (
		user        models.User
		passwdValue *string = nil
		IDTipo      *int    = nil
	)

	if err := c.Bind(&user); err != nil {
		log.Println(err)
		return c.JSON(http.StatusInternalServerError, "")
	}

	rows, err := s.db.Query(
		`SELECT CONTRASEÑA, IDTIPO FROM USUARIO WHERE NOMBREUSUARIO = ?`,
		user.Username,
	)
	if err != nil {
		log.Println(err)
		return c.JSON(http.StatusInternalServerError, "")
	}

	defer rows.Close()

	for rows.Next() {
		if err = rows.Scan(&passwdValue, &IDTipo); err != nil {
			log.Println(err)
			return c.JSON(http.StatusInternalServerError, "")
		}
	}

	if passwdValue == nil || IDTipo == nil {
		log.Println(err)
		return c.JSON(http.StatusBadRequest, errors.New("invalid user or password"))
	}

	if err = bcrypt.CompareHashAndPassword([]byte(*passwdValue), []byte(user.Password)); err != nil {
		log.Println(err)
		return c.JSON(http.StatusInternalServerError, "")
	}

	return c.JSON(http.StatusOK, models.User{
		Username: user.Username,
		IdTipo:   *IDTipo,
	})
}

func (s *Service) CreateUser(c echo.Context) error {
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
	INSERT INTO USUARIO( NOMBREUSUARIO, CONTRASEÑA, NOMBRE, APELLIDO, IDTIPO)
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

func (s *Service) CreateUserAccess(c echo.Context) error {
	var useraccess models.UserAccess

	if err := c.Bind(&useraccess); err != nil {
		log.Println(err)
		return c.JSON(http.StatusInternalServerError, "")
	}

	res, err := s.db.Exec("")
	if err != nil {
		log.Println(err)
		return c.JSON(http.StatusInternalServerError, "not created")
	}

	if num, err := res.RowsAffected(); err != nil || num < 1 {
		log.Println(err)
		return c.JSON(http.StatusInternalServerError, "not created")
	}

	return c.JSON(http.StatusOK, "Created successfully")

}

func (s *Service) DeleteUser(c echo.Context) error {
	nombreusuario := c.QueryParam(("username"))
	if nombreusuario == "" {
		log.Println(nombreusuario)
		return c.JSON(http.StatusBadRequest, "no user to delete")
	}

	res, err := s.db.Exec("DELETE FROM USUARIO WHERE NOMBREUSUARIO = ?", nombreusuario)
	if err != nil {
		log.Println(err)
		return c.JSON(http.StatusInternalServerError, "no deleted")
	}

	if num, err := res.RowsAffected(); err != nil || num < 1 {
		log.Println(err)
		return c.JSON(http.StatusInternalServerError, "no deleted")
	}

	return c.JSON(http.StatusOK, "Deleted successfully")
}

func hashAndSalt(pwd string) (string, error) {
	pwdByte := []byte(pwd)

	hash, err := bcrypt.GenerateFromPassword(pwdByte, bcrypt.MinCost)
	if err != nil {
		return "", err
	}

	return string(hash), nil
}
