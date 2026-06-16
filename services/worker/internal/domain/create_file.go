package domain

import (
	"io"
	"log"
	"mime/multipart"
	"os"
	"path/filepath"
)

func CreateFile(name string, data multipart.File) error {

	filename := filepath.Join("/services/worker/files_test", name)

	f, err := os.OpenFile(filename, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0644)
	if err != nil {
		log.Println("Error creating file:", err)
		return ErrCreatingFile
	}
	defer f.Close()

	byter, err := io.ReadAll(data)
	if err != nil {
		log.Println("Error reading request file:", err)
		return ErrReadingReq
	}

	if _, err = f.Write(byter); err != nil {
		log.Println("Error writing to file:", err)
		return ErrWritingToFile
	}
	return nil
}
