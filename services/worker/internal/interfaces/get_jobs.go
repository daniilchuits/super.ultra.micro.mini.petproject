package interfaces

import "worker/internal/domain"

type GetJobs interface {
	Get(user_id int) (*[]domain.Note, error)
}
