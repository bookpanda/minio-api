package dto

import "github.com/bookpanda/minio-api/internal/model"

type DecomposedFile struct {
	Filename string
	Data     []byte
}

type UploadFileRequest struct {
	Bucket string     `json:"bucket"`
	File   model.File `json:"file"`
}

type UploadFileResponse struct {
	Url string `json:"url"`
	Key string `json:"key"`
}

type DeleteFileRequest struct {
	Bucket string `json:"name" validate:"required,min=3,max=50"`
	FileId string `json:"fileId" validate:"required,min=3,max=50"`
}

type DeleteFileResponse struct {
	Success bool `json:"success"`
}

type GetFileRequest struct {
	Bucket string `json:"name" validate:"required,min=3,max=50"`
	FileId string `json:"fileId" validate:"required,min=3,max=50"`
}

type GetFileResponse struct {
	FileUrl string `json:"fileUrl"`
}
