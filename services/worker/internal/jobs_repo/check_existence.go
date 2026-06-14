package jobsrepo

func (repo repoManager) CheckExistence(user_id int, filename string) (bool, error) {

	query := `
		SELECT EXISTS(
			SELECT 1 
			FROM jobs 
			WHERE user_id=$1
				AND filename=$2
		)
	`

	var exists bool
	err := repo.db.QueryRow(query, user_id, filename).Scan(&exists)
	return exists, err
}
