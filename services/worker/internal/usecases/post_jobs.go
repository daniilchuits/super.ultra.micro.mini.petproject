package usecases

import (
	"log"
	"worker/internal/domain"
	"worker/internal/interfaces"
)

type PostJobUsecase struct {
	CheckExistence  interfaces.CheckExistence
	InsertInterface interfaces.PostNote
}

func (post PostJobUsecase) PostUsecase(user_id int, filename domain.Filename) (*domain.Note, error) {

	exists, _, err := post.CheckExistence.CheckExistence(user_id, filename.Name)

	if err != nil {
		log.Println("Err checking existence:", err)
		return nil, domain.ErrDuringCheckingExistence
	}

	if exists {
		return nil, domain.ErrNoteInPendingStatus
	}

	note, err := post.InsertInterface.Post(user_id, filename.Name)
	if err != nil {
		log.Println(err)
		return nil, domain.ErrDuringInsertingNote
	}
	return note, nil
}
