package usecases

import (
	"worker/internal/domain"
	"worker/internal/interfaces"
)

type GetJobs struct {
	GetInt interfaces.GetJobs
}

func (getUc GetJobs) GetJ(user_id int) (*[]domain.Note, error) {
	return getUc.GetInt.Get(user_id)
}
