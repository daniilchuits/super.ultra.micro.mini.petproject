package interfaces

type UpdateStatus interface {
	UpdateProc(note_id, user_id int) error
}
