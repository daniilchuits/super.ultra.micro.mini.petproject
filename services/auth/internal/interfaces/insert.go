package interfaces

import "auth/internal/domain"

type InsertNote interface {
	InsertIntoUsers(login, password string) (*domain.Credentials, error)
}
