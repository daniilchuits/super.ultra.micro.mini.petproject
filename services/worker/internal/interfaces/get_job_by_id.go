package interfaces

import "worker/internal/domain"

type GetJobInterface interface {
	GetJob(note_id, user_id int) (*domain.Note, error)
}
