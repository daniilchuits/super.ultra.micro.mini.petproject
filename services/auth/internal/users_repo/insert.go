package usersrepo

import (
	"auth/internal/domain"
	"log"
)

func (repo repoManager) InsertIntoUsers(login, password string) (*domain.Credentials, error) {

	query := `
	INSERT INTO users (login, password) VALUES
	($1,$2)
	RETURNING id
	`

	var id int

	err := repo.db.QueryRow(query, login, password).Scan(&id)
	if err != nil {
		log.Println(err)
		return nil, domain.ErrDuringInsertingToUsers
	}
	return &domain.Credentials{
		Id:       id,
		Login:    login,
		Password: password,
	}, nil
}
