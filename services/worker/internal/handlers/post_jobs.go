package handlers

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"strconv"
	"worker/internal/domain"
	"worker/internal/interfaces"
	"worker/internal/transport"
	"worker/internal/usecases"
)

type postJobHandler struct {
	uc usecases.PostJobUsecase
}

func NewPostJobHandler(
	check interfaces.CheckExistence,
	post interfaces.PostNote,
) *postJobHandler {
	return &postJobHandler{uc: usecases.PostJobUsecase{
		CheckExistence:  check,
		InsertInterface: post,
	}}
}

func (post *postJobHandler) PostNote(w http.ResponseWriter, r *http.Request) {

	userIDStr := r.Header.Get("user_id")
	userID, err := strconv.Atoi(userIDStr)
	if err != nil {
		log.Println(err)
		http.Error(w, "Error during converting user_id from jwt", 500)
		return
	}

	if userID == 0 {
		http.Error(w, "Server got wrong user_id from jwt", 500)
		return
	}

	var filenameStructHTTP transport.Filename
	if err = json.NewDecoder(r.Body).Decode(&filenameStructHTTP); err != nil {
		http.Error(w, "Error decoding filename", 400)
		return
	}

	filenameDomain := transport.ImportToDomain(filenameStructHTTP)
	hasSufix := domain.ValidateFilename(filenameDomain)
	if !hasSufix {
		http.Error(w, "File doesn't have sufix '.txt'", 400)
		return
	}

	noteDomain, err := post.uc.PostUsecase(userID, filenameDomain)
	if err != nil {

		if errors.Is(err, domain.ErrNoteInPendingStatus) {
			http.Error(w, domain.ErrNoteInPendingStatus.Error(), 400)
			return
		} else if errors.Is(err, domain.ErrDuringCheckingExistence) {
			http.Error(w, domain.ErrDuringCheckingExistence.Error(), 500)
			return
		} else if errors.Is(err, domain.ErrDuringInsertingNote) {
			http.Error(w, domain.ErrDuringInsertingNote.Error(), 500)
			return
		}

		log.Println(err)
		http.Error(w, "Unknown error", 500)
		return
	}

	note := transport.ImportToHttp(*noteDomain)

	if err = json.NewEncoder(w).Encode(&note); err != nil {
		http.Error(w, "Error encoding but it is inserted", 500)
		return
	}
}
