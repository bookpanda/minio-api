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
	Url     string `json:"url"`
	FileKey string `json:"fileKey"`
}

type DeleteFileRequestBody struct {
	FileKey string `json:"fileKey" validate:"required,min=3,max=50"`
}

type DeleteFileRequest struct {
	Bucket  string `json:"bucket"`
	FileKey string `json:"fileKey"`
}

type DeleteFileResponse struct {
	Success bool `json:"success"`
}

type GetFileRequest struct {
	Bucket  string `json:"name"`
	FileKey string `json:"fileKey"`
}

type GetFileResponse struct {
	FileUrl string `json:"fileUrl"`
}
