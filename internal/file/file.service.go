package file

import "go.uber.org/zap"

type Service interface {
	Upload()
	Delete()
	Get()
}

type serviceImpl struct {
	repo   Repository
	logger *zap.Logger
}

func NewService(repo Repository, logger *zap.Logger) Service {
	return &serviceImpl{
		repo:   repo,
		logger: logger,
	}
}

func (r *serviceImpl) Upload() {}

func (r *serviceImpl) Delete() {}

func (r *serviceImpl) Get() {}
