package model

type File struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	Data []byte `json:"data"`
}
