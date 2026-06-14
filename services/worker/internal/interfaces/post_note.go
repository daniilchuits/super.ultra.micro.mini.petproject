package interfaces

import "worker/internal/domain"

type PostNote interface {
	Post(user_id int, filename string) (*domain.Note, error)
}
