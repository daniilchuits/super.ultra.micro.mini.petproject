package usecases

import (
	"log"
	"worker/internal/domain"
	"worker/internal/interfaces"
)

type DoneUsecase struct {
	ChooseProcFile  interfaces.ChooseProcFile
	FinalUpdateNote interfaces.FinalUpdateNote
}

func (do DoneUsecase) Exec() (string, error) {

	noteID, filename, err := do.ChooseProcFile.Choose()
	if err != nil {
		log.Println("Error in getting 'processing' file")
		return "", err
	}

	data := domain.Process(noteID, filename)

	if err = do.FinalUpdateNote.Update(data); err != nil {
		log.Println("Error updating jobs table:", err)
		return filename, err
	}
	return filename, nil
}
