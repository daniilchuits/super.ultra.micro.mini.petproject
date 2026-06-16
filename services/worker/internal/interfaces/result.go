package interfaces

import "mime/multipart"

type InsertIntoResultsInterface interface {
	Insert(file multipart.File, filename string) error
}
