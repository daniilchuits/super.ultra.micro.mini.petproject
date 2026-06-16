package jobsrepo

import (
	"database/sql"
	"errors"
)

func (repo *repoManager) CheckExistence(user_id int, filename string) (bool, int, error) {

	query := `
		SELECT id 
		FROM jobs
		WHERE user_id=$1
			AND filename=$2
		LIMIT 1
	`

	var noteID int
	err := repo.db.QueryRow(query, user_id, filename).Scan(&noteID)

	switch {
	case errors.Is(err, sql.ErrNoRows):
		return false, 0, nil
	case err != nil:
		return true, 0, err
	default:
		return true, noteID, nil
	}
}
