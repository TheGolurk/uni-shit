package internal

import (
	"errors"
	"net/http"
	"time"

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
		return c.JSON(http.StatusInternalServerError, "")
	}

	rows, err := s.db.Query(
		`SELECT CONTRASEÑA, IDTIPO FROM USUARIO WHERE NOMBREUSUARIO = ?`,
		user.Username,
	)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, "")
	}

	defer rows.Close()

	for rows.Next() {
		if err = rows.Scan(&passwdValue, &IDTipo); err != nil {
			return c.JSON(http.StatusInternalServerError, "")
		}
	}

	if passwdValue == nil || IDTipo == nil {
		return c.JSON(http.StatusBadRequest, errors.New("invalid user or password"))
	}

	if err = bcrypt.CompareHashAndPassword([]byte(*passwdValue), []byte(user.Password)); err != nil {
		return c.JSON(http.StatusInternalServerError, "")
	}

	c.SetCookie(&http.Cookie{
		Name:    "login-user",
		Value:   user.Username,
		Expires: time.Now().Add(24 * time.Hour),
	})

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
		return c.JSON(http.StatusInternalServerError, "")
	}

	pwd, err := hashAndSalt(user.Password)
	if err != nil {
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
		return c.JSON(http.StatusInternalServerError, "")
	}

	return c.JSON(http.StatusOK, "created")
}

func hashAndSalt(pwd string) (string, error) {
	pwdByte := []byte(pwd)

	hash, err := bcrypt.GenerateFromPassword(pwdByte, bcrypt.MinCost)
	if err != nil {
		return "", err
	}

	return string(hash), nil
}
