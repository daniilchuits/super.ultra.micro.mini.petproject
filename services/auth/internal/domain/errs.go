package domain

import "errors"

var (
	// register
	ErrUserExists              error = errors.New("User already exists")
	ErrDuringCheckingExistance error = errors.New("Error during checking existance")
	ErrDuringInsertingToUsers  error = errors.New("Error during inserting to users table")
)
