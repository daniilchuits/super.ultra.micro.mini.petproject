package interfaces

type InterfaceToGetOneNote interface {
	GetOneNote(login string) (string, bool, error)
}
