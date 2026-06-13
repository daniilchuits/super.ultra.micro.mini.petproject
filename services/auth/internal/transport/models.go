package transport

type Credentials struct {
	Id       int    `json:"id"`
	Login    string `json:"login"`
	Password string `json:"password"`
}
