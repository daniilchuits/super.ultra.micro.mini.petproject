package usecases

import (
	"worker/internal/domain"
	"worker/internal/interfaces"
)

type GetJobUsecase struct {
	GetJob interfaces.GetJobInterface
}

func (gj GetJobUsecase) GetJobById(note_id, user_id int) (*domain.Note, error) {
	return gj.GetJobById(note_id, user_id)
}
