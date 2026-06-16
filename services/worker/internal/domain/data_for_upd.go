package domain

type ProcessedData struct {
	JobID        int
	Name         string
	Status       string
	LinesCount   *int
	WordsCount   *int
	CharsCount   *int
	ErrorMessage *string
}
