package file

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type Handler interface {
	Upload(c *gin.Context)
	Delete(c *gin.Context)
	Get(c *gin.Context)
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

func (r *handlerImpl) Upload(c *gin.Context) {}

func (r *handlerImpl) Delete(c *gin.Context) {}

func (r *handlerImpl) Get(c *gin.Context) {}
