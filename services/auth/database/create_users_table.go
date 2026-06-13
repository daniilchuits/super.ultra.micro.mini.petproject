package database

import "database/sql"

func CreateUsersTable(db *sql.DB) error {

	query := `
	CREATE TABLE IF NOT EXISTS users(
		id INT PRIMARY KEY GENERATED ALWAYS AS IDENTITY,
		login TEXT NOT NULL,
		password TEXT NOT NULL
	)
	`

	_, err := db.Exec(query)
	return err
}
