package jobsrepo

import (
	"log"
	"worker/internal/domain"
)

func (repo repoManager) Get(user_id int) (*[]domain.Note, error) {

	query := `
		SELECT * FROM jobs WHERE user_id=$1
	`

	rows, err := repo.db.Query(query, user_id)
	if err != nil {
		log.Println("Err during execing query:", err)
		return nil, domain.ErrScaningJobs
	}
	defer rows.Close()

	var notes []domain.Note

	for rows.Next() {

		var note domain.Note
		if err = rows.Scan(
			&note.Id,
			&note.UserId,
			&note.File,
			&note.Status,
			&note.CreatedAt,
		); err != nil {
			log.Println("Error during scaning one note:", err)
			return nil, domain.ErrScaningJobs
		}
		notes = append(notes, note)
	}

	if err = rows.Err(); err != nil {
		log.Println("Error after scaning:", err)
		return nil, domain.ErrScaningJobs
	}
	return &notes, nil
}
