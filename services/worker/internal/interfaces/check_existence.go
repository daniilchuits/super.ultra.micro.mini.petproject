package interfaces

type CheckExistence interface {
	CheckExistence(user_id int, filename string) (bool, error)
}
