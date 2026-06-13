package domain

type Credentials struct {
	Id       int
	Login    string
	Password string
}

type RegisteredData struct {
	Id       int
	Password string
}
