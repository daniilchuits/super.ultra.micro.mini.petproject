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

type registerHandler struct {
	uc usecases.UsecaseToRegister
}

func NewRegisterHandler(
	getNote interfaces.InterfaceToGetOneNote,
	insert interfaces.InsertNote,
) registerHandler {
	return registerHandler{uc: usecases.UsecaseToRegister{
		ExistsInUsers: getNote,
		InsertOneNote: insert,
	}}
}

func (reg registerHandler) RegisterUser(w http.ResponseWriter, r *http.Request) {

	var cred transport.Credentials
	if err := json.NewDecoder(r.Body).Decode(&cred); err != nil {
		log.Println(err)
		http.Error(w, "Error decoding credentials", 400)
		return
	}

	credDomain := transport.ImportToDomain(cred)

	newCred, err := reg.uc.RegisterUser(credDomain)
	if err != nil {

		if errors.Is(err, domain.ErrUserExists) {
			http.Error(w, domain.ErrUserExists.Error(), 400)
			return
		} else if errors.Is(err, domain.ErrDuringCheckingExistance) {
			http.Error(w, domain.ErrDuringCheckingExistance.Error(), 500)
			return
		} else if errors.Is(err, domain.ErrDuringInsertingToUsers) {
			http.Error(w, domain.ErrDuringInsertingToUsers.Error(), 500)
			return
		}

		log.Println(err)
		http.Error(w, "Unknown error", 500)
		return
	}

	httpCred := transport.ImportToHttp(*newCred)

	if err = json.NewEncoder(w).Encode(httpCred); err != nil {
		log.Println(err)
		http.Error(w, "Encoding error, but note inserted", 500)
		return
	}
}
