package domain

import "time"

type Filename struct {
	Name string
}

type Note struct {
	Id        int
	UserId    int
	File      string
	Status    string
	CreatedAt time.Time
}
