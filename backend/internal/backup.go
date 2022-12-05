package internal

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"net/http"
	"os/exec"
	"time"
)

func (s *Service) Backup(c echo.Context) error {
	t := time.Now()
	path := fmt.Sprintf("backups/%s", t.Format("2006-02-01-03:04"))
	_, err := exec.Command("cp", "-R", "mysqldata", path).Output()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, "")
	}

	return c.JSON(http.StatusCreated, "")
}
