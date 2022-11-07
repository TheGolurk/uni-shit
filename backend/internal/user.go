package internal

import (
	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"upemor.com/shit-project/models"
)

func (s *Service) Login(c echo.Context) error {
	var (
		user        models.User
		passwdValue *string = nil
	)

	if err := c.Bind(&user); err != nil {
		return err
	}

	rows, err := s.db.Query(
		`SELECT CONTRASEÑA FROM USUARIO WHERE NOMBREUSUARIO = ?`,
		user.Username,
	)
	if err != nil {
		return err
	}

	defer rows.Close()

	for rows.Next() {
		if err = rows.Scan(&passwdValue); err != nil {
			return err
		}
	}

	if err = bcrypt.CompareHashAndPassword([]byte(*passwdValue), []byte(user.Password)); err != nil {
		return err
	}

	return c.JSON(http.StatusOK, "login ok")
}

func (s *Service) CreateUser(c echo.Context) error {
	var (
		user models.User
	)

	if err := c.Bind(&user); err != nil {
		return err
	}

	pwd, err := hashAndSalt(user.Password)
	if err != nil {
		return err
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
		return err
	}

	return nil
}

func hashAndSalt(pwd string) (string, error) {
	pwdByte := []byte(pwd)

	hash, err := bcrypt.GenerateFromPassword(pwdByte, bcrypt.MinCost)
	if err != nil {
		return "", err
	}

	return string(hash), nil
}
