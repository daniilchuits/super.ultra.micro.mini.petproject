package transport

import "auth/internal/domain"

func ImportToHttp(cred domain.Credentials) Credentials {
	return Credentials{
		Id:       cred.Id,
		Login:    cred.Login,
		Password: cred.Password,
	}
}
