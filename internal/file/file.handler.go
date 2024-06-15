package file

import (
	"net/http"
	"strings"

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

// Upload godoc
// @Summary      Upload object to minio
// @Description  get string by ID
// @Tags         file
// @Accept       mpfd
// @Produce      json
// @Param        bucket formData string true "Which bucket to upload to"
// @Param        file formData file true "file to upload"
// @Success      200  {object}  dto.UploadFileResponse
// @Failure      400  {object}  errors.AppError
// @Failure      401  {object}  errors.AppError
// @Failure      500  {object}  errors.AppError
// @Router       /api/v1/upload [post]
func (h *handlerImpl) Upload(c router.Context) {
	bucket := c.PostForm("bucket")
	if bucket == "" {
		c.ResponseError(errors.BadRequestError("bucket is required"))
		return
	}

	file, err := c.FormFile("file", h.allowedContentType, h.maxFileSize)
	if err != nil {
		c.ResponseError(errors.BadRequestError(err.Error()))
		return
	}

	req := &dto.UploadFileRequest{
		Bucket: bucket,
		File: model.File{
			ID:   c.NewUUID(),
			Name: file.Filename,
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

func (h *handlerImpl) Get(c router.Context) {
	bucket := c.Param("bucket")
	if bucket == "" {
		c.ResponseError(errors.BadRequestError("bucket route parameter is required"))
		return
	}

	objectKey := c.Query("key")
	if objectKey == "" {
		c.ResponseError(errors.BadRequestError("key query parameter is required"))
		return
	}

	req := &dto.GetFileRequest{
		Bucket:  bucket,
		FileKey: objectKey,
	}

	res, err := h.svc.Get(req)
	if err != nil {
		c.ResponseError(errors.InternalServerError(err.Error()))
		return
	}

	c.JSON(http.StatusOK, res)
}

func (h *handlerImpl) Delete(c router.Context) {
	bucket := c.Param("bucket")
	if bucket == "" {
		c.ResponseError(errors.BadRequestError("bucket route parameter is required"))
		return
	}

	body := &dto.DeleteFileRequestBody{}
	if err := c.Bind(body); err != nil {
		c.ResponseError(errors.BadRequestError(err.Error()))
		return
	}

	if errorList := h.validate.Validate(body); errorList != nil {
		c.ResponseError(errors.BadRequestError(strings.Join(errorList, ", ")))
		return
	}

	req := &dto.DeleteFileRequest{
		Bucket:  bucket,
		FileKey: body.FileKey,
	}

	res, err := h.svc.Delete(req)
	if err != nil {
		c.ResponseError(errors.InternalServerError(err.Error()))
		return
	}

	c.JSON(http.StatusOK, res)
}
