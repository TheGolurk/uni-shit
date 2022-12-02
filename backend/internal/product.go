package internal

import (
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
	"upemor.com/shit-project/models"
)

func (s *Service) CreateProduct(c echo.Context) error {
	var (
		product models.Producto
	)

	if err := c.Bind(&product); err != nil {
		log.Println(err)
		return c.JSON(http.StatusBadRequest, "")
	}

	_, err := s.db.Exec(
		Producto_create,
		product.PesoProducto,
		product.PrecioProduccto,
	)

	if err != nil {
		log.Println(err)
		return c.JSON(http.StatusBadRequest, "")
	}

	return c.JSON(http.StatusOK, "created")
}

func (s *Service) DeleteProduct(c echo.Context) error {
	var (
		product models.Producto
	)

	if err := c.Bind(&product); err != nil {
		log.Println(err)
		return c.JSON(http.StatusBadRequest, "")
	}

	_, err := s.db.Exec(Producto_delete, product.IdProducto)

	if err != nil {
		log.Println(err)
		return c.JSON(http.StatusBadRequest, "")
	}

	return c.JSON(http.StatusOK, "deleted")

}

func (s *Service) UpdateProduct(c echo.Context) error {
	var (
		product models.Producto
	)

	if err := c.Bind(&product); err != nil {
		log.Println(err)
		return c.JSON(http.StatusBadRequest, "")
	}

	_, err := s.db.Exec(Producto_update, product.PesoProducto, product.PrecioProduccto, product.IdProducto)

	if err != nil {
		log.Println(err)
		return c.JSON(http.StatusBadRequest, "")
	}

	return c.JSON(http.StatusOK, "updated")
}

func (s *Service) GetProduct(c echo.Context) error {
	var (
		product  models.Producto
		products []models.Producto
	)

	rows, err := s.db.Query(Producto_select)
	if err != nil {
		log.Println(err)
		return c.JSON(http.StatusBadRequest, "")
	}

	defer rows.Close()

	for rows.Next() {
		if err = rows.Scan(
			&product.PesoProducto,
			&product.PrecioProduccto,
		); err != nil {

			log.Println(err)
			return c.JSON(http.StatusInternalServerError, "")
		}
		products = append(products, product)
	}

	return c.JSON(http.StatusOK, products)
}
