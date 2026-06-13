package usersrepo

import (
	"auth/internal/domain"
	"log"
)

func (repo repoManager) GetOneNote(login string) (*domain.RegisteredData, bool, error) {

	query := `
		SELECT EXISTS (
			SELECT 1 FROM users WHERE login=$1
		)
	`

	var exists bool
	if err := repo.db.QueryRow(query, login).Scan(&exists); err != nil {
		return nil, false, err
	}
	if !exists {
		return nil, false, nil
	}

	query = `
	SELECT id, password FROM users WHERE login=$1
	`

	var (
		id       int
		password string
	)

	err := repo.db.QueryRow(query, login).Scan(&id, &password)
	if err != nil {
		log.Println(err)
		return nil, true, err
	}
	return &domain.RegisteredData{
		Id:       id,
		Password: password,
	}, true, err
}
