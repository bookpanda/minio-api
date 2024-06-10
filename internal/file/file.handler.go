package file

import (
	"net/http"
	"strings"

	"github.com/bookpanda/minio-api/errors"
	"github.com/bookpanda/minio-api/internal/dto"
	"github.com/bookpanda/minio-api/internal/validator"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type Handler interface {
	Upload(c *gin.Context)
	Delete(c *gin.Context)
	Get(c *gin.Context)
}

type handlerImpl struct {
	svc      Service
	validate validator.DtoValidator
	log      *zap.Logger
}

func NewHandler(svc Service, validate validator.DtoValidator, log *zap.Logger) Handler {
	return &handlerImpl{
		svc:      svc,
		validate: validate,
		log:      log,
	}
}

func (h *handlerImpl) Upload(c *gin.Context) {
	req := &dto.UploadFileRequest{}
	if err := c.BindJSON(req); err != nil {
		errors.ResponseError(c, errors.BadRequest)
		return
	}

	if errorList := h.validate.Validate(req); errorList != nil {
		errors.ResponseError(c, errors.BadRequestError(strings.Join(errorList, ", ")))
		return
	}

	res, err := h.svc.Upload(c, *req)
	if err != nil {
		errors.ResponseError(c, errors.InternalServerError(err.Error()))
		return
	}

	c.JSON(http.StatusOK, res)
}

func (h *handlerImpl) Delete(c *gin.Context) {}

func (h *handlerImpl) Get(c *gin.Context) {}
