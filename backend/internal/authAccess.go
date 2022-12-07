package internal

import "errors"

func (s *Service) AuthUserAccess(id, tabla string) error {
	count := 0
	rows, err := s.db.Query("SELECT COUNT(*) FROM ACCESO WHERE ID_TIPO = ? AND TABLAS = ?;", id, tabla)
	if err != nil {
		return err
	}

	defer rows.Close()

	for rows.Next() {
		if err = rows.Scan(&count); err != nil {
			return err
		}
	}

	if count == 0 {
		return errors.New("not auth")
	}

	return nil
}
