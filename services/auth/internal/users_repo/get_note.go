package usersrepo

import "log"

func (repo repoManager) GetOneNote(login string) (string, bool, error) {

	query := `
		SELECT EXISTS (
			SELECT 1 FROM users WHERE login=$1
		)
	`

	var exists bool
	if err := repo.db.QueryRow(query, login).Scan(&exists); err != nil {
		return "", false, err
	}
	if !exists {
		return "", false, nil
	}

	query = `
	SELECT password FROM users WHERE login=$1
	`

	var password string

	err := repo.db.QueryRow(query, login).Scan(&password)
	if err != nil {
		log.Println(err)
		return "", true, err
	}
	return password, true, err
}
