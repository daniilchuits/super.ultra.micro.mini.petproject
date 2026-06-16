package jobsrepo

import "worker/internal/domain"

func (repo repoManager) UpdateProc(note_id, user_id int) error {

	query := `
		UPDATE jobs
		SET status=$1
		WHERE id=$2
			AND user_id=$3
	`

	_, err := repo.db.Exec(query, domain.ProcStatus, note_id, user_id)
	return err
}
