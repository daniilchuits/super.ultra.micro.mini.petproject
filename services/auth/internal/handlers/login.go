package handlers

import (
	"auth/internal/domain"
	"auth/internal/interfaces"
	"auth/internal/transport"
	"auth/internal/usecases"
	"encoding/json"
	"errors"
	"log"
	"net/http"
)

type loginHandler struct {
	uc        usecases.LoginUsecase
	secretKey string
}

func NewLoginHandler(
	getNote interfaces.InterfaceToGetOneNote,
	secretKey string,
) loginHandler {
	return loginHandler{
		uc: usecases.LoginUsecase{
			CheckPassword: getNote,
		},
		secretKey: secretKey,
	}
}

func (login loginHandler) LoginUser(w http.ResponseWriter, r *http.Request) {

	var cred transport.Credentials
	if err := json.NewDecoder(r.Body).Decode(&cred); err != nil {
		log.Println(r.Body)
		http.Error(w, "Error decoding credentials", 400)
		return
	}

	domainCred := transport.ImportToDomain(cred)

	jwt, err := login.uc.LoginUser(domainCred.Login, domainCred.Password, login.secretKey)
	if err != nil {

		if errors.Is(err, domain.ErrUserIsNotRegisteres) {
			http.Error(w, domain.ErrUserIsNotRegisteres.Error(), 400)
			return
		} else if errors.Is(err, domain.ErrWrongPassword) {
			http.Error(w, domain.ErrWrongPassword.Error(), 400)
			return
		} else if errors.Is(err, domain.ErrMakingJWT) {
			http.Error(w, domain.ErrMakingJWT.Error(), 500)
			return
		}

		log.Println(err)
		http.Error(w, "Unknown error", 500)
		return
	}

	jwtHTTP := transport.NewJWTTransport(jwt)
	if err = json.NewEncoder(w).Encode(jwtHTTP); err != nil {
		log.Println(err)
		http.Error(w, "Error encoding jwt", 500)
		return
	}
}
