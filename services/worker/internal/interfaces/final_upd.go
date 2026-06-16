package interfaces

import "worker/internal/domain"

type FinalUpdateNote interface {
	Update(data *domain.ProcessedData) error
}
