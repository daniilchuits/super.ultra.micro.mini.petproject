package jobsrepo

import (
	"time"
	"worker/internal/domain"
)

func (repo repoManager) Update(data *domain.ProcessedData) error {

	query := `
	UPDATE jobs
	SET 
		status=$1,
		lines_count=$2,
		words_count=$3,
		chars_count=$4,
		error_mesage=$5,
		processed_at=$6
	WHERE id=$7
		AND filename=$8
	RETURNING
		id, 
		user_id,
		filename,
		status,
		lines_count,
		words_count,
		chars_count,
		error_mesage,
		created_at,
		processed_at
	`

	if _, err := repo.db.Exec(
		query,
		domain.DoneStatus,
		data.LinesCount,
		data.WordsCount,
		data.CharsCount,
		data.ErrorMessage,
		time.Now(),
		data.JobID,
		data.Name,
	); err != nil {
		return domain.ErrDuringUpdNote
	}

	return nil
}
