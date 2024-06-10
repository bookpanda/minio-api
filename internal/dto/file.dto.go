package dto

import "github.com/bookpanda/minio-api/internal/model"

type UploadFileRequest struct {
	Bucket string     `json:"name"`
	File   model.File `json:"file"`
}

type UploadFileResponse struct {
	Url string `json:"url"`
	Key string `json:"key"`
}

type DeleteFileRequest struct {
	Bucket string `json:"name"`
	FileId string `json:"fileId"`
}

type DeleteFileResponse struct {
	Success bool `json:"success"`
}

type GetFileRequest struct {
	Bucket string `json:"name"`
	FileId string `json:"fileId"`
}

type GetFileResponse struct {
	FileUrl string `json:"fileUrl"`
}
