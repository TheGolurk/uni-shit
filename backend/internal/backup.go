package internal

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"net/http"
	"os/exec"
	"time"
)

func (s *Service) Backup(c echo.Context) error {
	_, err := s.db.Query(
		`INSERT INTO BACKUP_HISTORY(DATE_OF_BACKUP, WHO) VALUES(CURRENT_DATE, ?);`,
		c.QueryParam("username"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, "")
	}

	t := time.Now()
	path := fmt.Sprintf("backups/%s", t.Format("2006-02-01-03:04"))
	_, err = exec.Command("cp", "-R", "mysqldata", path).Output()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, "")
	}

	return c.JSON(http.StatusCreated, "")
}
