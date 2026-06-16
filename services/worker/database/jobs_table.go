package database

import "database/sql"

type dbManager struct {
	db *sql.DB
}

func NewDbManager(db *sql.DB) dbManager {
	return dbManager{db: db}
}

func (manager dbManager) CreateJobsTable() error {

	query := `
		CREATE TABLE IF NOT EXISTS jobs(
			id INT PRIMARY KEY GENERATED ALWAYS AS IDENTITY,
			user_id INT,
			filename TEXT NOT NULL,
			status TEXT NOT NULL,
			lines_count INT,
			words_count INT,
			chars_count INT,
			error_mesage TEXT,
			created_at TIMESTAMP DEFAULT now(),
			processed_at TIMESTAMP
		)
	`

	// to late i saw that wrote 'error_mesage' :D so much to update, so pofig

	_, err := manager.db.Exec(query)
	return err
}
