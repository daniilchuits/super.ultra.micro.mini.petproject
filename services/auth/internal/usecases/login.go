package usecases

import (
	"auth/internal/domain"
	"auth/internal/interfaces"
	"log"
)

type LoginUsecase struct {
	CheckPassword interfaces.InterfaceToGetOneNote
}

func (uc LoginUsecase) LoginUser(login, password, secret string) (string, error) {

	registeredData, exists, err := uc.CheckPassword.GetOneNote(login)
	if err != nil {
		log.Println(err)
		return "", err
	}
	if !exists {
		return "", domain.ErrUserIsNotRegisteres
	}

	if err = domain.CompareWithHashedPassword(registeredData.Password, password); err != nil {
		return "", domain.ErrWrongPassword
	}

	jwt, err := domain.MakingJWT(registeredData.Id, secret)
	if err != nil {
		log.Println(err)
		return "", domain.ErrMakingJWT
	}
	return jwt, nil
}
