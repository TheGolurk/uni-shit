package internal

import (
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
	"upemor.com/shit-project/models"
)

func (s *Service) CreateSell(c echo.Context) error {
	var (
		sell models.Venta
	)

	if err := c.Bind(&sell); err != nil {
		log.Println(err)
		return c.JSON(http.StatusInternalServerError, "")
	}

	_, err := s.db.Exec(Venta_create, sell.IdUsuarioVenta, sell.IdPro, sell.Total,
		sell.Iva, sell.FechaVenta, sell.IdCli)

	if err != nil {
		log.Println(err)
		return c.JSON(http.StatusInternalServerError, "")
	}

	return c.JSON(http.StatusOK, "created")
}

func (s *Service) DeleteSell(c echo.Context) error {
	var (
		venta models.Venta
	)

	if err := c.Bind(&venta); err != nil {
		log.Println(err)
		return c.JSON(http.StatusInternalServerError, "")
	}

	_, err := s.db.Exec(Venta_delete, venta.IdVenta)

	if err != nil {
		log.Println(err)
		return c.JSON(http.StatusInternalServerError, "")
	}

	return c.JSON(http.StatusOK, "deleted")
}

func (s *Service) UpdateSell(c echo.Context) error {
	var (
		venta models.Venta
	)

	if err := c.Bind(&venta); err != nil {
		log.Println(err)
		return c.JSON(http.StatusInternalServerError, "")
	}

	_, err := s.db.Exec(Venta_update)

	if err != nil {
		log.Println(err)
		return c.JSON(http.StatusInternalServerError, "")
	}

	return c.JSON(http.StatusOK, "updated")
}

func (s *Service) GetSell(c echo.Context) error {
	var (
		venta  models.Venta
		ventas []models.Venta
	)

	rows, err := s.db.Query(Venta_select)
	if err != nil {
		log.Println(err)
		return c.JSON(http.StatusInternalServerError, "")
	}

	defer rows.Close()

	for rows.Next() {
		if err = rows.Scan(
			&venta.IdUsuarioVenta,
			&venta.IdPro,
			&venta.Total,
			&venta.Iva,
			&venta.FechaVenta,
			&venta.IdCli,
		); err != nil {
			log.Println(err)
			return c.JSON(http.StatusInternalServerError, "")
		}
		ventas = append(ventas, venta)
	}

	return c.JSON(http.StatusOK, ventas)
}
