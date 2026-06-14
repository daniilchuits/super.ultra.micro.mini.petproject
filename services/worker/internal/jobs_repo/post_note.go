package jobsrepo

import "worker/internal/domain"

func (repo repoManager) Post(user_id int, filename string) (*domain.Note, error) {

	query := `
		INSERT INTO jobs (user_id, filename, status) VALUES
		($1,$2,$3)
		RETURNING *
	`

	var note domain.Note

	err := repo.db.QueryRow(query, user_id, filename, domain.PendStatus).Scan(
		&note.Id,
		&note.UserId,
		&note.File,
		&note.Status,
		&note.CreatedAt,
	)
	return &note, err
}
