package usecases

import (
	"auth/internal/domain"
	"auth/internal/interfaces"
	"log"
)

type LoginUsecase struct {
	GetNote interfaces.InterfaceToGetOneNote
}

func (uc LoginUsecase) LoginUser(credentials domain.Credentials, secret string) (string, error) {

	if err := domain.Validate(credentials); err != nil {
		return "", err
	}

	registeredData, exists, err := uc.GetNote.GetOneNote(credentials.Login)
	if err != nil {
		log.Println(err)
		return "", err
	}
	if !exists {
		return "", domain.ErrUserIsNotRegisteres
	}

	if err = domain.CompareWithHashedPassword(registeredData.Password, credentials.Password); err != nil {
		return "", domain.ErrWrongPassword
	}

	jwt, err := domain.MakingJWT(registeredData.Id, secret)
	if err != nil {
		log.Println(err)
		return "", domain.ErrMakingJWT
	}
	return jwt, nil
}
