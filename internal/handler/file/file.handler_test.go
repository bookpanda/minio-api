package file

import (
	"errors"
	"net/http"
	"testing"

	"github.com/bookpanda/minio-api/apperrors"
	"github.com/bookpanda/minio-api/constants"
	"github.com/bookpanda/minio-api/internal/dto"
	"github.com/bookpanda/minio-api/internal/model"
	mock_metrics "github.com/bookpanda/minio-api/mocks/metrics"
	mock_context "github.com/bookpanda/minio-api/mocks/router"
	mock_file "github.com/bookpanda/minio-api/mocks/service"
	mock_validator "github.com/bookpanda/minio-api/mocks/validator"
	"github.com/golang/mock/gomock"
	"github.com/google/uuid"
	"github.com/stretchr/testify/suite"
	"go.uber.org/zap"
)

type FileHandlerTest struct {
	suite.Suite
	controller         *gomock.Controller
	logger             *zap.Logger
	maxFileSize        int64
	allowedContentType map[string]struct{}
}

func TestFileHandler(t *testing.T) {
	suite.Run(t, new(FileHandlerTest))
}

func (t *FileHandlerTest) SetupTest() {
	t.controller = gomock.NewController(t.T())
	t.logger = zap.NewNop()
	t.maxFileSize = 10
	t.allowedContentType = map[string]struct{}{}
}

func (t *FileHandlerTest) TestUploadSuccess() {
	svc := mock_file.NewMockService(t.controller)
	ctx := mock_context.NewMockContext(t.controller)
	metrics := mock_metrics.NewMockRequestsMetrics(t.controller)
	id := uuid.New()
	uploadFileResponse := &dto.UploadFileResponse{
		Url:     "url",
		FileKey: "object",
	}

	ctx.EXPECT().PostForm("bucket").Return("bucket")
	ctx.EXPECT().FormFile("file", t.allowedContentType, t.maxFileSize).Return(&dto.DecomposedFile{
		Filename: "object",
		Data:     []byte("data"),
	}, nil)
	ctx.EXPECT().NewUUID().Return(id)
	svc.EXPECT().Upload(&dto.UploadFileRequest{
		File: model.File{
			ID:   id,
			Name: "object",
			Data: []byte("data"),
		},
		Bucket: "bucket",
	}).Return(uploadFileResponse, nil)
	metrics.EXPECT().AddRequest(constants.File, constants.POST, http.StatusOK)
	ctx.EXPECT().JSON(http.StatusOK, uploadFileResponse)

	hdr := NewHandler(svc, nil, t.maxFileSize, t.allowedContentType, t.logger, metrics)
	hdr.Upload(ctx)
}

func (t *FileHandlerTest) TestUploadNoBucket() {
	ctx := mock_context.NewMockContext(t.controller)
	metrics := mock_metrics.NewMockRequestsMetrics(t.controller)

	ctx.EXPECT().PostForm("bucket").Return("")
	metrics.EXPECT().AddRequest(constants.File, constants.POST, http.StatusBadRequest)
	ctx.EXPECT().ResponseError(apperrors.BadRequestError("bucket is required"))

	hdr := NewHandler(nil, nil, t.maxFileSize, t.allowedContentType, t.logger, metrics)
	hdr.Upload(ctx)
}

func (t *FileHandlerTest) TestUploadFormFileError() {
	ctx := mock_context.NewMockContext(t.controller)
	metrics := mock_metrics.NewMockRequestsMetrics(t.controller)

	ctx.EXPECT().PostForm("bucket").Return("bucket")
	ctx.EXPECT().FormFile("file", t.allowedContentType, t.maxFileSize).Return(nil, errors.New("error1"))
	metrics.EXPECT().AddRequest(constants.File, constants.POST, http.StatusBadRequest)
	ctx.EXPECT().ResponseError(apperrors.BadRequestError("error1"))

	hdr := NewHandler(nil, nil, t.maxFileSize, t.allowedContentType, t.logger, metrics)
	hdr.Upload(ctx)
}

func (t *FileHandlerTest) TestUploadServiceError() {
	svc := mock_file.NewMockService(t.controller)
	ctx := mock_context.NewMockContext(t.controller)
	metrics := mock_metrics.NewMockRequestsMetrics(t.controller)

	ctx.EXPECT().PostForm("bucket").Return("bucket")
	ctx.EXPECT().FormFile("file", t.allowedContentType, t.maxFileSize).Return(&dto.DecomposedFile{
		Filename: "object",
		Data:     []byte("data"),
	}, nil)
	ctx.EXPECT().NewUUID().Return(uuid.New())
	svc.EXPECT().Upload(gomock.Any()).Return(nil, &apperrors.AppError{})
	metrics.EXPECT().AddRequest(constants.File, constants.POST, apperrors.AppError{}.HttpCode)
	ctx.EXPECT().ResponseError(&apperrors.AppError{})

	hdr := NewHandler(svc, nil, t.maxFileSize, t.allowedContentType, t.logger, metrics)
	hdr.Upload(ctx)
}

func (t *FileHandlerTest) TestGetSuccess() {
	svc := mock_file.NewMockService(t.controller)
	ctx := mock_context.NewMockContext(t.controller)
	metrics := mock_metrics.NewMockRequestsMetrics(t.controller)
	getFileResponse := &dto.GetFileResponse{
		FileUrl: "url",
	}

	ctx.EXPECT().Param("bucket").Return("bucket")
	ctx.EXPECT().Query("key").Return("object")
	svc.EXPECT().Get(&dto.GetFileRequest{
		Bucket:  "bucket",
		FileKey: "object",
	}).Return(getFileResponse, nil)
	metrics.EXPECT().AddRequest(constants.File, constants.GET, http.StatusOK)
	ctx.EXPECT().JSON(http.StatusOK, getFileResponse)

	hdr := NewHandler(svc, nil, t.maxFileSize, t.allowedContentType, t.logger, metrics)
	hdr.Get(ctx)
}

func (t *FileHandlerTest) TestGetNoBucket() {
	ctx := mock_context.NewMockContext(t.controller)
	metrics := mock_metrics.NewMockRequestsMetrics(t.controller)

	ctx.EXPECT().Param("bucket").Return("")
	metrics.EXPECT().AddRequest(constants.File, constants.GET, http.StatusBadRequest)
	ctx.EXPECT().ResponseError(apperrors.BadRequestError("bucket route parameter is required"))

	hdr := NewHandler(nil, nil, t.maxFileSize, t.allowedContentType, t.logger, metrics)
	hdr.Get(ctx)
}

func (t *FileHandlerTest) TestGetNoKey() {
	ctx := mock_context.NewMockContext(t.controller)
	metrics := mock_metrics.NewMockRequestsMetrics(t.controller)

	ctx.EXPECT().Param("bucket").Return("bucket")
	ctx.EXPECT().Query("key").Return("")
	metrics.EXPECT().AddRequest(constants.File, constants.GET, http.StatusBadRequest)
	ctx.EXPECT().ResponseError(apperrors.BadRequestError("key query parameter is required"))

	hdr := NewHandler(nil, nil, t.maxFileSize, t.allowedContentType, t.logger, metrics)
	hdr.Get(ctx)
}

func (t *FileHandlerTest) TestGetServiceError() {
	svc := mock_file.NewMockService(t.controller)
	ctx := mock_context.NewMockContext(t.controller)
	metrics := mock_metrics.NewMockRequestsMetrics(t.controller)

	ctx.EXPECT().Param("bucket").Return("bucket")
	ctx.EXPECT().Query("key").Return("object")
	svc.EXPECT().Get(gomock.Any()).Return(nil, &apperrors.AppError{})
	metrics.EXPECT().AddRequest(constants.File, constants.GET, apperrors.AppError{}.HttpCode)
	ctx.EXPECT().ResponseError(&apperrors.AppError{})

	hdr := NewHandler(svc, nil, t.maxFileSize, t.allowedContentType, t.logger, metrics)
	hdr.Get(ctx)
}

func (t *FileHandlerTest) TestDeleteSuccess() {
	svc := mock_file.NewMockService(t.controller)
	ctx := mock_context.NewMockContext(t.controller)
	validator := mock_validator.NewMockDtoValidator(t.controller)
	metrics := mock_metrics.NewMockRequestsMetrics(t.controller)
	deleteFileResponse := &dto.DeleteFileResponse{
		Success: true,
	}
	body := &dto.DeleteFileRequestBody{}

	ctx.EXPECT().Param("bucket").Return("bucket")
	ctx.EXPECT().Bind(body).Return(nil)
	validator.EXPECT().Validate(body).Return(nil)
	svc.EXPECT().Delete(&dto.DeleteFileRequest{
		Bucket:  "bucket",
		FileKey: body.FileKey,
	}).Return(deleteFileResponse, nil)
	metrics.EXPECT().AddRequest(constants.File, constants.DELETE, http.StatusOK)
	ctx.EXPECT().JSON(http.StatusOK, deleteFileResponse)

	hdr := NewHandler(svc, validator, t.maxFileSize, t.allowedContentType, t.logger, metrics)
	hdr.Delete(ctx)
}

func (t *FileHandlerTest) TestDeleteNoBucket() {
	ctx := mock_context.NewMockContext(t.controller)
	metrics := mock_metrics.NewMockRequestsMetrics(t.controller)

	ctx.EXPECT().Param("bucket").Return("")
	metrics.EXPECT().AddRequest(constants.File, constants.DELETE, http.StatusBadRequest)
	ctx.EXPECT().ResponseError(apperrors.BadRequestError("bucket route parameter is required"))

	hdr := NewHandler(nil, nil, t.maxFileSize, t.allowedContentType, t.logger, metrics)
	hdr.Delete(ctx)
}

func (t *FileHandlerTest) TestDeleteBindError() {
	ctx := mock_context.NewMockContext(t.controller)
	metrics := mock_metrics.NewMockRequestsMetrics(t.controller)
	body := &dto.DeleteFileRequestBody{}

	ctx.EXPECT().Param("bucket").Return("bucket")
	ctx.EXPECT().Bind(body).Return(errors.New("error1"))
	metrics.EXPECT().AddRequest(constants.File, constants.DELETE, http.StatusBadRequest)
	ctx.EXPECT().ResponseError(apperrors.BadRequestError("error1"))

	hdr := NewHandler(nil, nil, t.maxFileSize, t.allowedContentType, t.logger, metrics)
	hdr.Delete(ctx)
}

func (t *FileHandlerTest) TestDeleteValidationErrors() {
	ctx := mock_context.NewMockContext(t.controller)
	validator := mock_validator.NewMockDtoValidator(t.controller)
	metrics := mock_metrics.NewMockRequestsMetrics(t.controller)
	body := &dto.DeleteFileRequestBody{}

	ctx.EXPECT().Param("bucket").Return("bucket")
	ctx.EXPECT().Bind(body).Return(nil)
	validator.EXPECT().Validate(body).Return([]string{"error1", "error2"})
	metrics.EXPECT().AddRequest(constants.File, constants.DELETE, http.StatusBadRequest)
	ctx.EXPECT().ResponseError(apperrors.BadRequestError("error1, error2"))

	hdr := NewHandler(nil, validator, t.maxFileSize, t.allowedContentType, t.logger, metrics)
	hdr.Delete(ctx)
}

func (t *FileHandlerTest) TestDeleteServiceError() {
	svc := mock_file.NewMockService(t.controller)
	ctx := mock_context.NewMockContext(t.controller)
	validator := mock_validator.NewMockDtoValidator(t.controller)
	metrics := mock_metrics.NewMockRequestsMetrics(t.controller)
	body := &dto.DeleteFileRequestBody{}

	ctx.EXPECT().Param("bucket").Return("bucket")
	ctx.EXPECT().Bind(body).Return(nil)
	validator.EXPECT().Validate(body).Return(nil)
	svc.EXPECT().Delete(&dto.DeleteFileRequest{
		Bucket:  "bucket",
		FileKey: body.FileKey,
	}).Return(nil, &apperrors.AppError{})
	metrics.EXPECT().AddRequest(constants.File, constants.DELETE, apperrors.AppError{}.HttpCode)
	ctx.EXPECT().ResponseError(&apperrors.AppError{})

	hdr := NewHandler(svc, validator, t.maxFileSize, t.allowedContentType, t.logger, metrics)
	hdr.Delete(ctx)
}
