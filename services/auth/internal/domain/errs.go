package domain

import "errors"

var (
	// register
	ErrUserExists              error = errors.New("User already exists")
	ErrDuringCheckingExistance error = errors.New("Error during checking existance")
	ErrDuringInsertingToUsers  error = errors.New("Error during inserting to users table")

	// login
	ErrUserIsNotRegisteres error = errors.New("User not found")
	ErrWrongPassword       error = errors.New("Wrong password")
	ErrMakingJWT           error = errors.New("Error making jwt")

	// makingHashedPassword
	ErrMakingHashedPassword error = errors.New("Error making HashedPassword")

	// validate
	ErrShortLogin    error = errors.New("Short login, must be at least 6 symbols")
	ErrShortPassword error = errors.New("Short password, must be at least 6 symbols")
)
