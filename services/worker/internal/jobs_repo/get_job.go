package jobsrepo

import (
	"log"
	"worker/internal/domain"
)

func (repo *repoManager) GetJob(note_id, user_id int) (*domain.Note, error) {

	query := `
	SELECT 
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
	FROM jobs
	WHERE id=$1
		AND user_id=$2
	`

	var note domain.Note

	err := repo.db.QueryRow(query, note_id, user_id).Scan(
		&note.Id,
		&note.UserId,
		&note.File,
		&note.Status,
		&note.LinesCount,
		&note.WordsCount,
		&note.CharsCount,
		&note.ErrorMessage,
		&note.CreatedAt,
		&note.ProcessedAt,
	)

	if err != nil {
		log.Println("Scanning one note err:", err)
		return nil, domain.ErrScaningOneJob
	}
	return &note, nil
}
