package transport

import "auth/internal/domain"

func ImportToDomain(cred Credentials) domain.Credentials {
	return domain.Credentials{
		Login:    cred.Login,
		Password: cred.Password,
	}
}
