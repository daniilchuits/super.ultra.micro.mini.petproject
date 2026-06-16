package usecases

import (
	"log"
	"mime/multipart"
	"os"
	"worker/internal/domain"
	"worker/internal/interfaces"
)

type UpdateJobUsecase struct {
	Check              interfaces.CheckExistence
	UpdateStatusToProc interfaces.UpdateStatus
	Get                interfaces.GetJobInterface
}

func (upd UpdateJobUsecase) Exec(user_id int, filename string, fileData multipart.File) (*domain.Note, error) {

	exists, note_id, err := upd.Check.CheckExistence(user_id, filename)
	if err != nil {
		log.Println("Checkint existence err:", err)
		return nil, domain.ErrCheckingExistence
	}

	if !exists {
		return nil, domain.ErrNoteDoesnotExists
	}

	if err = upd.UpdateStatusToProc.UpdateProc(note_id, user_id); err != nil {
		log.Println("Upd to 'proc' err:", err)
		return nil, domain.ErrUpdatingToProc
	}

	wd, _ := os.Getwd()
	log.Println("WD:", wd)

	log.Println("Creating file:", filename)

	if err = domain.CreateFile(filename, fileData); err != nil {
		return nil, err
	}

	note, err := upd.Get.GetJob(note_id, user_id)
	if err != nil {
		return nil, err
	}

	return note, nil
}
