package goworker

import (
	"log"
	"time"
	"worker/internal/interfaces"
	"worker/internal/usecases"
)

type goWorker struct {
	uc usecases.DoneUsecase
}

func NewGoWorker(
	choose interfaces.ChooseProcFile,
	upd interfaces.FinalUpdateNote,
) *goWorker {
	return &goWorker{
		uc: usecases.DoneUsecase{
			ChooseProcFile:  choose,
			FinalUpdateNote: upd,
		},
	}
}

func (work *goWorker) Process() {

	for {

		time.Sleep(5 * time.Second)

		if filename, err := work.uc.Exec(); err != nil {
			log.Printf("Worker coudn't proc file %s: %v\n", filename, err)
		}
	}
}
