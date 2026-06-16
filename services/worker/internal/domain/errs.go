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

	//POST update
	ErrNoteDoesnotExists error = errors.New("Note doesn't exists")
	ErrCheckingExistence error = errors.New("Error during checking existence")
	ErrUpdatingToProc    error = errors.New("Error during updating status to 'processing'")
	ErrGettingFile       error = errors.New("Error getting file")

	// CREATING FILE
	ErrCreatingFile  error = errors.New("Error creating file")
	ErrReadingReq    error = errors.New("Error reading request file")
	ErrWritingToFile error = errors.New("Error writing to file")

	// WORKER
	ErrInsertingIntoResultsTable error = errors.New("Insert into results table err")
	ErrUpdatingToDone            error = errors.New("Error during updating status to 'done'")

	// GoWorker
	ErrNoFileInDir   error = errors.New("No file in 'files_test'")
	ErrReadingFile   error = errors.New("Error reading file")
	ErrDuringUpdNote error = errors.New("Error during updating note")
	ErrNoProcFiles   error = errors.New("There is no 'processing' files")
)
