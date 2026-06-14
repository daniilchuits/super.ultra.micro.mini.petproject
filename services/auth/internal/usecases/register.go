package usecases

import (
	"auth/internal/domain"
	"auth/internal/interfaces"
)

type UsecaseToRegister struct {
	ExistsInUsers interfaces.InterfaceToGetOneNote
	InsertOneNote interfaces.InsertNote
}

func (uc UsecaseToRegister) RegisterUser(credentials domain.Credentials) (*domain.Credentials, error) {

	if err := domain.Validate(credentials); err != nil {
		return nil, err
	}

	_, exists, err := uc.ExistsInUsers.GetOneNote(credentials.Login)
	if err != nil {
		return nil, domain.ErrDuringCheckingExistance
	}
	if exists {
		return nil, domain.ErrUserExists
	}

	hashedPassword, err := domain.CreateHashedPassword(credentials.Password)
	if err != nil {
		return nil, domain.ErrMakingHashedPassword
	}

	return uc.InsertOneNote.InsertIntoUsers(credentials.Login, hashedPassword)
}
