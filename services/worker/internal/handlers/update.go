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

type updateHandler struct {
	uc usecases.UpdateJobUsecase
}

func NewUpdateHandler(
	check interfaces.CheckExistence,
	update interfaces.UpdateStatus,
	get interfaces.GetJobInterface,
) *updateHandler {
	return &updateHandler{uc: usecases.UpdateJobUsecase{
		Check:              check,
		UpdateStatusToProc: update,
		Get:                get,
	}}
}

func (upd *updateHandler) UpdateFile(w http.ResponseWriter, r *http.Request) {

	file, header, err := r.FormFile("file")
	if err != nil {
		log.Println("Getting file err:", err)
		http.Error(w, domain.ErrGettingFile.Error(), 500)
		return
	}

	userIdStr := r.Header.Get("user_id")
	id, err := strconv.Atoi(userIdStr)
	if err != nil {
		log.Println("Converting 'user_id' to int err:", err)
		http.Error(w, "Error during converting user_id from jwt", 500)
		return
	}

	noteDomain, err := upd.uc.Exec(id, header.Filename, file)
	if err != nil {

		if errors.Is(err, domain.ErrCreatingFile) {
			http.Error(w, domain.ErrCreatingFile.Error(), 500)
			return
		} else if errors.Is(err, domain.ErrReadingReq) {
			http.Error(w, domain.ErrReadingReq.Error(), 500)
			return
		} else if errors.Is(err, domain.ErrWritingToFile) {
			http.Error(w, domain.ErrWritingToFile.Error(), 500)
		} else {

			log.Println("UC err:", err)
			http.Error(w, "Unknown err", 500)
			return
		}
	}

	if err = json.NewEncoder(w).Encode(transport.ImportToHttp(*noteDomain)); err != nil {
		log.Println("Error encoding but job uploaded:", err)
		http.Error(w, "Error encoding but job uploaded", 500)
		return
	}
}
