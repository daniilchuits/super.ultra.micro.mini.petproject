package domain

import "errors"

var (
	ErrEmptyAuthorization      error = errors.New("Empty authorization header")
	ErrNoBearerInAuthorization error = errors.New("Authorization header doesn't have 'Bearer'")
	ErrTokenIsNotValid         error = errors.New("Token is not already valid")
	ErrParsingToken            error = errors.New("Error validating token")
	ErrInvalidUserId           error = errors.New("Not valid 'user_id' in claims")
)
