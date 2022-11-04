package main

import (
	"database/sql"
	"net/http"

	"github.com/labstack/echo/v4"
)

type config struct {
	db *sql.DB
}

type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func main() {
	db, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/ESTANCIA")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	e := echo.New()

	ccc := config{
		db: db,
	}

	e.POST("/login", ccc.Login)

	e.Logger.Fatal(e.Start(":8080"))

}

func (cc *config) Login(c echo.Context) error {
	var (
		user string
		a    User
	)

	c.Bind(&a)

	rows, err := cc.db.Query("SELECT NOMBREUSUARIO FROM USUARIO WHERE NOMBREUSUARIO = ? AND CONTRASEÑA = ?", a.Username, a.Password)
	if err != nil {
		return c.JSON(http.StatusBadRequest, "login err")
	}

	defer rows.Close()

	for rows.Next() {
		if err = rows.Scan(&user); err != nil {
			return c.JSON(http.StatusInternalServerError, "login err")
		}
	}

	return c.JSON(http.StatusOK, "login ok")
}
