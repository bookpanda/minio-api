package file

import (
	"fmt"
	"strings"

	"github.com/bookpanda/minio-api/apperrors"
	"github.com/bookpanda/minio-api/internal/dto"
	"go.uber.org/zap"
)

type Service interface {
	Upload(req *dto.UploadFileRequest) (res *dto.UploadFileResponse, apperr *apperrors.AppError)
	Delete(req *dto.DeleteFileRequest) (res *dto.DeleteFileResponse, apperr *apperrors.AppError)
	Get(req *dto.GetFileRequest) (res *dto.GetFileResponse, apperr *apperrors.AppError)
}

type serviceImpl struct {
	repo Repository
	log  *zap.Logger
}

func NewService(repo Repository, log *zap.Logger) Service {
	return &serviceImpl{
		repo: repo,
		log:  log,
	}
}

func (s *serviceImpl) Upload(req *dto.UploadFileRequest) (res *dto.UploadFileResponse, apperr *apperrors.AppError) {
	objectKey := req.File.ID.String()[:8] + "_" + strings.ReplaceAll(req.File.Name, " ", "_")

	url, key, err := s.repo.Upload(req.File.Data, req.Bucket, objectKey)
	if err != nil {
		s.log.Named("file svc").Error("Couldn't upload object", zap.Error(err))
		return nil, apperrors.InternalServerError(fmt.Sprintf("Couldn't upload object to %v/%v.", req.Bucket, objectKey))
	}

	return &dto.UploadFileResponse{
		Url:     url,
		FileKey: key,
	}, nil
}

func (s *serviceImpl) Delete(req *dto.DeleteFileRequest) (res *dto.DeleteFileResponse, apperr *apperrors.AppError) {
	err := s.repo.Delete(req.Bucket, req.FileKey)
	if err != nil {
		s.log.Named("file svc").Error("Couldn't delete object", zap.Error(err))
		return nil, apperrors.InternalServerError(fmt.Sprintf("Couldn't delete object %v/%v.", req.Bucket, req.FileKey))
	}

	return &dto.DeleteFileResponse{
		Success: true,
	}, nil
}

func (s *serviceImpl) Get(req *dto.GetFileRequest) (res *dto.GetFileResponse, apperr *apperrors.AppError) {
	url, err := s.repo.Get(req.Bucket, req.FileKey)
	if err != nil {
		s.log.Named("file svc").Error("Couldn't get object", zap.Error(err))
		return nil, apperrors.InternalServerError(fmt.Sprintf("Couldn't get object %v/%v.", req.Bucket, req.FileKey))
	}
	if url == "" {
		return nil, apperrors.NotFoundError(fmt.Sprintf("Couldn't find object %v/%v.", req.Bucket, req.FileKey))
	}

	return &dto.GetFileResponse{
		FileUrl: url,
	}, nil
}
