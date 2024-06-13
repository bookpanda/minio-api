package model

import "github.com/google/uuid"

type File struct {
	ID   uuid.UUID `json:"id"`
	Name string    `json:"name"`
	Data []byte    `json:"data"`
}
