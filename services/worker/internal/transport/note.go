package transport

import "time"

type Note struct {
	Id        int       `json:"id"`
	UserId    int       `json:"user_id"`
	File      string    `json:"filename"`
	Status    string    `json:"status"`
	CreatedAt time.Time `json:"created_at"`
}
