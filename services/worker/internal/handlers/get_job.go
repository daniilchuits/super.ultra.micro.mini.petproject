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

	"github.com/go-chi/chi/v5"
)

type getOneNoteHandler struct {
	uc usecases.GetJobUsecase
}

func NewGetOneNoteHandler(getJob interfaces.GetJobInterface) getOneNoteHandler {
	return getOneNoteHandler{uc: usecases.GetJobUsecase{
		GetJob: getJob,
	}}
}

func (get getOneNoteHandler) GetOneJob(w http.ResponseWriter, r *http.Request) {

	noteIdStr := chi.URLParam(r, "id")
	noteId, err := strconv.Atoi(noteIdStr)
	if err != nil {
		http.Error(w, "Error converting 'id' from URL to int", 400)
		return
	}

	userIdStr := r.Header.Get("user_id")
	userId, err := strconv.Atoi(userIdStr)
	if err != nil {
		http.Error(w, "Error converting 'user_id' from request", 400)
		return
	}

	note, err := get.uc.GetJobById(noteId, userId)
	if err != nil {

		if errors.Is(err, domain.ErrScaningOneJob) {
			http.Error(w, domain.ErrScaningOneJob.Error(), 500)
			return
		}

		log.Println("Handler:", err)
		http.Error(w, "Unknown error", 500)
		return
	}

	httpNote := transport.ImportToHttp(*note)

	if err = json.NewEncoder(w).Encode(httpNote); err != nil {
		log.Println(err)
		http.Error(w, "Error encoding", 500)
		return
	}
}
