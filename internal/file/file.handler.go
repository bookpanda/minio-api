package file

import "go.uber.org/zap"

type Handler interface {
	Upload()
	Delete()
	Get()
}

type handlerImpl struct {
	svc    Service
	logger *zap.Logger
}

func NewHandler(svc Service, logger *zap.Logger) Handler {
	return &handlerImpl{
		svc:    svc,
		logger: logger,
	}
}

func (r *handlerImpl) Upload() {}

func (r *handlerImpl) Delete() {}

func (r *handlerImpl) Get() {}
