package file

import (
	"net/http"

	"github.com/bookpanda/minio-api/constants"
	"github.com/bookpanda/minio-api/errors"
	"github.com/bookpanda/minio-api/internal/dto"
	"github.com/bookpanda/minio-api/internal/model"
	"github.com/bookpanda/minio-api/internal/utils"
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
	svc         Service
	validate    validator.DtoValidator
	maxFileSize int64
	log         *zap.Logger
}

func NewHandler(svc Service, validate validator.DtoValidator, maxFileSize int64, log *zap.Logger) Handler {
	return &handlerImpl{
		svc:         svc,
		validate:    validate,
		maxFileSize: maxFileSize,
		log:         log,
	}
}

func (h *handlerImpl) Upload(c *gin.Context) {
	bucket := c.PostForm("bucket")
	if bucket == "" {
		errors.ResponseError(c, errors.BadRequestError("bucket is required"))
		return
	}

	name := c.PostForm("name")
	if name == "" {
		errors.ResponseError(c, errors.BadRequestError("name is required"))
		return
	}

	file, err := c.FormFile("file")
	if err != nil {
		errors.ResponseError(c, errors.BadRequest)
		return
	}

	data, err := utils.ExtractFile(file, constants.AllowContentType, h.maxFileSize)
	if err != nil {
		errors.ResponseError(c, errors.InternalServerError(err.Error()))
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
			Name: name,
			Data: data,
		},
	}

	res, err := h.svc.Upload(c, req)
	if err != nil {
		errors.ResponseError(c, errors.InternalServerError(err.Error()))
		return
	}

	c.JSON(http.StatusOK, res)
}

func (h *handlerImpl) Delete(c *gin.Context) {}

func (h *handlerImpl) Get(c *gin.Context) {}
