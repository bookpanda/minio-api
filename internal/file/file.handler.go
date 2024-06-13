package file

import (
	"net/http"

	"github.com/bookpanda/minio-api/errors"
	"github.com/bookpanda/minio-api/internal/dto"
	"github.com/bookpanda/minio-api/internal/model"
	"github.com/bookpanda/minio-api/internal/router"
	"github.com/bookpanda/minio-api/internal/validator"
	"go.uber.org/zap"
)

type Handler interface {
	Upload(c router.Context)
	Delete(c router.Context)
	Get(c router.Context)
}

type handlerImpl struct {
	svc                Service
	validate           validator.DtoValidator
	maxFileSize        int64
	allowedContentType map[string]struct{}
	log                *zap.Logger
}

func NewHandler(svc Service, validate validator.DtoValidator, maxFileSize int64, allowedContentType map[string]struct{}, log *zap.Logger) Handler {
	return &handlerImpl{
		svc:                svc,
		validate:           validate,
		maxFileSize:        maxFileSize,
		allowedContentType: allowedContentType,
		log:                log,
	}
}

func (h *handlerImpl) Upload(c router.Context) {
	bucket := c.PostForm("bucket")
	if bucket == "" {
		c.ResponseError(errors.BadRequestError("bucket is required"))
		return
	}

	name := c.PostForm("name")
	if name == "" {
		c.ResponseError(errors.BadRequestError("name is required"))
		return
	}

	file, err := c.FormFile("file", h.allowedContentType, h.maxFileSize)
	if err != nil {
		c.ResponseError(errors.BadRequestError(err.Error()))
		return
	}

	// if err := c.BindJSON(req); err != nil {
	// 	errors.ResponseError(c, errors.BadRequest)
	// 	return
	// }

	// if errorList := h.validate.Validate(req); errorList != nil {
	// 	errors.ResponseError(c, errors.BadRequestError(strings.Join(errorList, ", ")))
	// 	return
	// }
	req := &dto.UploadFileRequest{
		Bucket: bucket,
		File: model.File{
			ID:   c.NewUUID(),
			Name: name,
			Data: file.Data,
		},
	}

	res, err := h.svc.Upload(req)
	if err != nil {
		c.ResponseError(errors.InternalServerError(err.Error()))
		return
	}

	c.JSON(http.StatusOK, res)
}

func (h *handlerImpl) Delete(c router.Context) {}

func (h *handlerImpl) Get(c router.Context) {}
