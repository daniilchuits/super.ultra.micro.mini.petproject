package jobsrepo

import (
	"database/sql"
	"errors"
	"log"
	"worker/internal/domain"
)

func (repo repoManager) Choose() (int, string, error) {

	query := `
	SELECT id, filename
	FROM jobs
	WHERE status=$1
	ORDER BY random()
	LIMIT 1
	`

	var (
		id       int
		filename string
	)

	err := repo.db.QueryRow(query, domain.ProcStatus).Scan(&id, &filename)
	if err != nil {
		log.Println("Cant select 'processing' file:", err)
		if errors.Is(err, sql.ErrNoRows) {
			return 0, "", domain.ErrNoProcFiles
		}
		return 0, "", err
	}
	return id, filename, nil
}
