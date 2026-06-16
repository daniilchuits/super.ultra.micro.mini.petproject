package interfaces

type ChooseProcFile interface {
	Choose() (int, string, error)
}
