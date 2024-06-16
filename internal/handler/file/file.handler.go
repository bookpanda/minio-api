package file

import (
	"net/http"
	"strings"

	"github.com/bookpanda/minio-api/apperrors"
	"github.com/bookpanda/minio-api/internal/dto"
	"github.com/bookpanda/minio-api/internal/model"
	"github.com/bookpanda/minio-api/internal/router"
	"github.com/bookpanda/minio-api/internal/service/file"
	"github.com/bookpanda/minio-api/internal/validator"
	"go.uber.org/zap"
)

type Handler interface {
	Upload(c router.Context)
	Delete(c router.Context)
	Get(c router.Context)
}

type handlerImpl struct {
	svc                file.Service
	validate           validator.DtoValidator
	maxFileSize        int64
	allowedContentType map[string]struct{}
	log                *zap.Logger
}

func NewHandler(svc file.Service, validate validator.DtoValidator, maxFileSize int64, allowedContentType map[string]struct{}, log *zap.Logger) Handler {
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
// @Description  specify the bucket and file to upload
// @Tags         file
// @Accept       mpfd
// @Produce      json
// @Param        bucket formData string true "Which bucket to upload to"
// @Param        file formData file true "file to upload"
// @Security     BearerAuth
// @Success      200  {object}  dto.UploadFileResponse
// @Failure      400  {object}  apperrors.AppError
// @Failure      401  {object}  apperrors.AppError
// @Failure      500  {object}  apperrors.AppError
// @Router       /file/upload [post]
func (h *handlerImpl) Upload(c router.Context) {
	bucket := c.PostForm("bucket")
	if bucket == "" {
		h.log.Named("file hdr").Error("bucket is required")
		c.ResponseError(apperrors.BadRequestError("bucket is required"))
		return
	}

	file, err := c.FormFile("file", h.allowedContentType, h.maxFileSize)
	if err != nil {
		h.log.Named("file hdr").Error("failed to parse form file", zap.Error(err))
		c.ResponseError(apperrors.BadRequestError(err.Error()))
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

	res, apperr := h.svc.Upload(req)
	if apperr != nil {
		h.log.Named("file hdr").Error("failed to upload file to service", zap.Error(apperr))
		c.ResponseError(apperr)
		return
	}

	c.JSON(http.StatusOK, res)
}

// Get godoc
// @Summary      Get object url from minio
// @Description  specify the bucket and object key to get
// @Tags         file
// @Accept       plain
// @Produce      json
// @Param        bucket path string true "Which bucket to get object from"
// @Param        key query string true "object key to get"
// @Security     BearerAuth
// @Success      200  {object}  dto.GetFileResponse
// @Failure      400  {object}  apperrors.AppError
// @Failure      401  {object}  apperrors.AppError
// @Failure      500  {object}  apperrors.AppError
// @Router       /file/get/{bucket} [get]
func (h *handlerImpl) Get(c router.Context) {
	bucket := c.Param("bucket")
	if bucket == "" {
		h.log.Named("file hdr").Error("bucket is required")
		c.ResponseError(apperrors.BadRequestError("bucket route parameter is required"))
		return
	}

	objectKey := c.Query("key")
	if objectKey == "" {
		h.log.Named("file hdr").Error("key query parameter is required")
		c.ResponseError(apperrors.BadRequestError("key query parameter is required"))
		return
	}

	req := &dto.GetFileRequest{
		Bucket:  bucket,
		FileKey: objectKey,
	}

	res, apperr := h.svc.Get(req)
	if apperr != nil {
		h.log.Named("file hdr").Error("failed to get file from service", zap.Error(apperr))
		c.ResponseError(apperr)
		return
	}

	c.JSON(http.StatusOK, res)
}

// Delete godoc
// @Summary      Delete object from minio
// @Description  specify the object's bucket and key to delete
// @Tags         file
// @Accept       json
// @Produce      json
// @Param        bucket path string true "Which bucket to delete object"
// @Param        fileKey body dto.DeleteFileRequestBody true "object key to delete"
// @Security     BearerAuth
// @Success      200  {object}  dto.DeleteFileResponse
// @Failure      400  {object}  apperrors.AppError
// @Failure      401  {object}  apperrors.AppError
// @Failure      500  {object}  apperrors.AppError
// @Router       /file/delete/{bucket} [delete]
func (h *handlerImpl) Delete(c router.Context) {
	bucket := c.Param("bucket")
	if bucket == "" {
		h.log.Named("file hdr").Error("bucket is required")
		c.ResponseError(apperrors.BadRequestError("bucket route parameter is required"))
		return
	}

	body := &dto.DeleteFileRequestBody{}
	if err := c.Bind(body); err != nil {
		h.log.Named("file hdr").Error("failed to bind request body", zap.Error(err))
		c.ResponseError(apperrors.BadRequestError(err.Error()))
		return
	}

	if errorList := h.validate.Validate(body); errorList != nil {
		h.log.Named("file hdr").Error("validation error", zap.Strings("errorList", errorList))
		c.ResponseError(apperrors.BadRequestError(strings.Join(errorList, ", ")))
		return
	}

	req := &dto.DeleteFileRequest{
		Bucket:  bucket,
		FileKey: body.FileKey,
	}

	res, apperr := h.svc.Delete(req)
	if apperr != nil {
		h.log.Named("file hdr").Error("failed to delete file from service", zap.Error(apperr))
		c.ResponseError(apperr)
		return
	}

	c.JSON(http.StatusOK, res)
}
