package transport

import "worker/internal/domain"

func ImportToHttp(note domain.Note) Note {
	return Note{
		Id:           note.Id,
		UserId:       note.UserId,
		File:         note.File,
		Status:       note.Status,
		LinesCount:   note.LinesCount,
		WordsCount:   note.WordsCount,
		CharsCount:   note.CharsCount,
		ErrorMessage: note.ErrorMessage,
		CreatedAt:    note.CreatedAt,
		ProcessedAt:  note.ProcessedAt,
	}
}
