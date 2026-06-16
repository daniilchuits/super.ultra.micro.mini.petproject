package transport

import "time"

type Note struct {
	Id           int        `json:"id"`
	UserId       int        `json:"user_id"`
	File         string     `json:"filename"`
	Status       string     `json:"status"`
	LinesCount   *int       `json:"lines_count,omitempty"`
	WordsCount   *int       `json:"words_count,omitempty"`
	CharsCount   *int       `json:"chars_count,omitempty"`
	ErrorMessage *string    `json:"error_message,omitempty"`
	CreatedAt    time.Time  `json:"created_at"`
	ProcessedAt  *time.Time `json:"processed_at,omitempty"`
}
