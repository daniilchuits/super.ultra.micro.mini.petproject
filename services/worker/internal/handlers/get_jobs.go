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

type getJobsHandler struct {
	uc usecases.GetJobs
}

func NewGetJobHandler(get interfaces.GetJobs) getJobsHandler {
	return getJobsHandler{uc: usecases.GetJobs{
		GetInt: get,
	}}
}

func (get getJobsHandler) GetJobs(w http.ResponseWriter, r *http.Request) {

	userIdStr := r.Header.Get("user_id")
	userId, err := strconv.Atoi(userIdStr)
	if err != nil {
		http.Error(w, "Error converting 'user_id' from request", 500)
		return
	}

	notesDomain, err := get.uc.GetJ(userId)
	if err != nil {

		if errors.Is(err, domain.ErrScaningJobs) {
			http.Error(w, domain.ErrScaningJobs.Error(), 500)
			return
		}

		log.Println(err)
		http.Error(w, "Unknown error", 500)
		return
	}

	var notes []transport.Note
	for _, note := range *notesDomain {
		notes = append(notes, transport.ImportToHttp(note))
	}

	if err = json.NewEncoder(w).Encode(notes); err != nil {
		http.Error(w, "Error encoding jobs", 500)
		return
	}
}
