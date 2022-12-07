package internal

import (
	"errors"
	"log"
	"net/http"
	"strconv"
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
		log.Println(err)
		return c.JSON(http.StatusInternalServerError, "")
	}

	rows, err := s.db.Query(
		`SELECT PASS, TIPO_ID FROM USUARIO WHERE USERNAME = ?`,
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

	if IDTipo != nil && *IDTipo != 1 {
		hourIniSTR := ""
		hourFinSTR := ""
		// get the hour
		rows, err := s.db.Query("SELECT HORARIO_INICIO, HORARIO_FINAL FROM TIPO_ACCESO WHERE ID_TIPO = ? LIMIT 1;", *IDTipo)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, "")
		}

		defer rows.Close()

		for rows.Next() {
			if err = rows.Scan(&hourIniSTR, &hourFinSTR); err != nil {
				return c.JSON(http.StatusInternalServerError, "")
			}
		}

		hour1, err := strconv.Atoi(hourIniSTR[:2])
		if err != nil {
			return c.JSON(http.StatusInternalServerError, "")
		}
		hour2, err := strconv.Atoi(hourFinSTR[:2])
		if err != nil {
			return c.JSON(http.StatusInternalServerError, "")
		}

		serverHour := time.Now().Hour()

		if serverHour < hour1 || serverHour > hour2 {
			return c.JSON(http.StatusUnauthorized, "")
		}
	}

	return c.JSON(http.StatusOK, models.User{
		Username: user.Username,
		IdTipo:   *IDTipo,
	})
}

func (s *Service) CreateUser(c echo.Context) error {
	var (
		user         models.User
		prevUsername string
	)

	if err := c.Bind(&user); err != nil {
		log.Println(err)
		return c.JSON(http.StatusBadRequest, "informacion erronea")
	}

	rows, err := s.db.Query("SELECT USERNAME FROM USUARIO WHERE USERNAME = ?", user.Username)
	if err != nil {
		log.Println(err)
		return c.JSON(http.StatusInternalServerError, "")
	}

	defer rows.Close()

	for rows.Next() {
		if err = rows.Scan(&prevUsername); err != nil {
			log.Println(err)
			return c.JSON(http.StatusInternalServerError, "")
		}
	}

	if prevUsername == user.Username {
		return c.JSON(http.StatusConflict, "duplicado")
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
		return c.JSON(http.StatusBadRequest, "")
	}

	return c.JSON(http.StatusOK, "created")
}

func (s *Service) DeleteUser(c echo.Context) error {
	nombreusuario := c.QueryParam(("username"))
	if nombreusuario == "" {
		log.Println(nombreusuario)
		return c.JSON(http.StatusBadRequest, "no user to delete")
	}

	res, err := s.db.Exec("DELETE FROM USUARIO WHERE USERNAME= ?", nombreusuario)
	if err != nil {
		log.Println(err)
		return c.JSON(http.StatusBadRequest, "no deleted")
	}

	if num, err := res.RowsAffected(); err != nil || num < 1 {
		log.Println(err)
		return c.JSON(http.StatusInternalServerError, "no deleted")
	}

	return c.JSON(http.StatusOK, "Deleted successfully")
}

func (s *Service) ModifyUser(c echo.Context) error {
	var (
		user models.User
	)

	if err := c.Bind(&user); err != nil {
		log.Println(err)
		return c.JSON(http.StatusBadRequest, "")
	}

	res, err := s.db.Exec("UPDATE USUARIO SET NOMBRE = ?, APELLIDO = ?, TIPO_ID = ? WHERE USERNAME = ?",
		user.Nombre, user.Apellido, user.IdTipo, user.Username)
	if err != nil {
		log.Println(err)
		return c.JSON(http.StatusBadRequest, "")
	}

	if count, err := res.RowsAffected(); count == 0 || err != nil {
		log.Println(err)
		return c.JSON(http.StatusInternalServerError, "")
	}

	return c.JSON(http.StatusOK, "modified successfully")
}

func hashAndSalt(pwd string) (string, error) {
	pwdByte := []byte(pwd)

	hash, err := bcrypt.GenerateFromPassword(pwdByte, bcrypt.MinCost)
	if err != nil {
		return "", err
	}

	return string(hash), nil
}

func (s *Service) GetUser(c echo.Context) error {
	var (
		user  models.User
		users []models.User
	)
	rows, err := s.db.Query("SELECT USERNAME, NOMBRE, APELLIDO, TIPO_ID FROM USUARIO;")
	if err != nil {
		log.Println(err)
		return c.JSON(http.StatusInternalServerError, "")
	}

	defer rows.Close()

	for rows.Next() {
		if err = rows.Scan(&user.Username, &user.Nombre, &user.Apellido, &user.IdTipo); err != nil {
			log.Println(err)
			return c.JSON(http.StatusInternalServerError, "")
		}
		users = append(users, user)
	}

	return c.JSON(http.StatusOK, users)
}
