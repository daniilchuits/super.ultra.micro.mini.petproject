package domain

import (
	"io"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func Process(noteID int, filename string) *ProcessedData {

	fullFilename := filepath.Join("files_test", filename)

	f, err := os.Open(fullFilename)
	if err != nil {
		log.Printf("Error opening file %s: %v\n", filename, err)
		msg := ErrNoFileInDir.Error()
		return &ProcessedData{
			JobID:        noteID,
			Name:         filename,
			Status:       string(FailedStatus),
			ErrorMessage: &msg,
		}
	}
	defer f.Close()

	dataBytes, err := io.ReadAll(f)
	if err != nil {
		log.Printf("Enable to read %s: %v\n", filename, err)
		msg := ErrReadingFile.Error()
		return &ProcessedData{
			JobID:        noteID,
			Name:         filename,
			Status:       string(FailedStatus),
			ErrorMessage: &msg,
		}
	}

	data := string(dataBytes)
	linesCount := strings.Count(data, "\n")
	wordsCount := len(strings.Fields(data))
	charsCount := len(data)

	return &ProcessedData{
		JobID:      noteID,
		Name:       filename,
		Status:     string(DoneStatus),
		LinesCount: &linesCount,
		WordsCount: &wordsCount,
		CharsCount: &charsCount,
	}
}
