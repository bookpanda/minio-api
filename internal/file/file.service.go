package file

import (
	"context"

	"github.com/bookpanda/minio-api/internal/dto"
	"go.uber.org/zap"
)

type Service interface {
	Upload(ctx context.Context, req dto.UploadFileRequest) (res *dto.UploadFileResponse, err error)
	Delete(ctx context.Context, req dto.DeleteFileRequest) (res *dto.DeleteFileResponse, err error)
	Get(ctx context.Context, req dto.GetFileRequest) (res *dto.GetFileResponse, err error)
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

func (s *serviceImpl) Upload(ctx context.Context, req dto.UploadFileRequest) (res *dto.UploadFileResponse, err error) {
	url, key, err := s.repo.Upload(req.File.Data, req.Bucket, req.File.ID)
	if err != nil {
		s.log.Named("file svc").Error("Couldn't upload object", zap.Error(err))
		return nil, err
	}

	return &dto.UploadFileResponse{
		Url: url,
		Key: key,
	}, nil
}

func (s *serviceImpl) Delete(ctx context.Context, req dto.DeleteFileRequest) (res *dto.DeleteFileResponse, err error) {
	err = s.repo.Delete(req.Bucket, req.FileId)
	if err != nil {
		s.log.Named("file svc").Error("Couldn't delete object", zap.Error(err))
		return nil, err
	}

	return &dto.DeleteFileResponse{
		Success: true,
	}, nil
}

func (s *serviceImpl) Get(ctx context.Context, req dto.GetFileRequest) (res *dto.GetFileResponse, err error) {
	url, err := s.repo.Get(req.Bucket, req.FileId)
	if err != nil {
		s.log.Named("file svc").Error("Couldn't get object", zap.Error(err))
		return nil, err
	}

	return &dto.GetFileResponse{
		FileUrl: url,
	}, nil
}
