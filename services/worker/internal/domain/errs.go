package domain

import "errors"

var (
	// POST jobs
	ErrNoteInPendingStatus     error = errors.New("File is already in pending status")
	ErrDuringCheckingExistence error = errors.New("Error during checking existence")
	ErrDuringInsertingNote     error = errors.New("Error during inserting note")

	// GET jobs
	ErrScaningJobs error = errors.New("Error scaning jobs")

	// GET jobs/{id}
	ErrScaningOneJob error = errors.New("Error scaning note")
)
