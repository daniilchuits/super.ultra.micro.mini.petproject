package transport

import "worker/internal/domain"

func ImportToHttp(note domain.Note) Note {
	return Note{
		Id:        note.Id,
		UserId:    note.UserId,
		File:      note.File,
		Status:    note.Status,
		CreatedAt: note.CreatedAt,
	}
}
