package domain

import "time"

type Filename struct {
	Name string
}

type Note struct {
	Id           int
	UserId       int
	File         string
	Status       string
	LinesCount   *int
	WordsCount   *int
	CharsCount   *int
	ErrorMessage *string
	CreatedAt    time.Time
	ProcessedAt  *time.Time
}
