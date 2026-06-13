package interfaces

import "auth/internal/domain"

type InterfaceToGetOneNote interface {
	GetOneNote(login string) (*domain.RegisteredData, bool, error)
}
