package internal

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"upemor.com/shit-project/models"
)

func (s *Service) Reportv1(c echo.Context) error {
	var (
		ccount  = 0
		tacceso = 0
		acceso  = 0
	)

	rows, err := s.db.Query(Report_Client)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, "")
	}

	defer rows.Close()

	for rows.Next() {
		if err = rows.Scan(&ccount); err != nil {
			return c.JSON(http.StatusInternalServerError, "")
		}
	}

	rows2, err := s.db.Query(Report_TipoAccess)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, "")
	}

	defer rows2.Close()

	for rows2.Next() {
		if err = rows2.Scan(&tacceso); err != nil {
			return c.JSON(http.StatusInternalServerError, "")
		}
	}

	rows3, err := s.db.Query(Report_Acess)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, "")
	}

	defer rows3.Close()

	for rows3.Next() {
		if err = rows3.Scan(&acceso); err != nil {
			return c.JSON(http.StatusInternalServerError, "")
		}
	}

	return c.JSON(http.StatusOK, models.R1{
		Cliente:     ccount,
		TipoAccesso: tacceso,
		Acceso:      acceso,
	})
}

func (s *Service) Reportv2(c echo.Context) error {
	var (
		total float64
	)

	rows, err := s.db.Query(Report_Sell)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, "")
	}

	defer rows.Close()

	for rows.Next() {
		if err = rows.Scan(&total); err != nil {
			return c.JSON(http.StatusInternalServerError, "")
		}
	}

	return c.JSON(http.StatusOK, models.R2{Total: total})
}

func (s *Service) Reportv3(c echo.Context) error {
	var (
		total int
	)
	who := c.QueryParam("who")
	where := c.QueryParam("where")

	rows, err := s.db.Query(Report_db, where, who)
	if err != nil {
		return c.JSON(http.StatusBadRequest, "")
	}

	defer rows.Close()

	for rows.Next() {
		if err = rows.Scan(&total); err != nil {
			return c.JSON(http.StatusBadRequest, "")
		}
	}
	if total == 0 {
		return c.JSON(http.StatusOK, models.R3{Total: 0})
	}

	return c.JSON(http.StatusOK, models.R3{Total: total})
}
