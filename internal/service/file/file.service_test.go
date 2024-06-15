package file

import (
	"errors"
	"testing"

	"github.com/bookpanda/minio-api/internal/dto"
	"github.com/bookpanda/minio-api/internal/model"
	mock_file "github.com/bookpanda/minio-api/mocks/repository"
	"github.com/golang/mock/gomock"
	"github.com/google/uuid"
	"github.com/stretchr/testify/suite"
	"go.uber.org/zap"
)

type FileServiceTest struct {
	suite.Suite
	controller *gomock.Controller
	logger     *zap.Logger
}

func TestFileService(t *testing.T) {
	suite.Run(t, new(FileServiceTest))
}

func (t *FileServiceTest) SetupTest() {
	t.controller = gomock.NewController(t.T())
	t.logger = zap.NewNop()
}

func (t *FileServiceTest) TestUploadSuccess() {
	id := uuid.New()
	objectKey := id.String()[:8] + "_" + "object"

	repo := mock_file.NewMockRepository(t.controller)
	repo.EXPECT().Upload([]byte("data"), "bucket", objectKey).Return("url", "object", nil)

	svc := NewService(repo, t.logger)

	res, err := svc.Upload(&dto.UploadFileRequest{
		File: model.File{
			ID:   id,
			Name: "object",
			Data: []byte("data"),
		},
		Bucket: "bucket",
	})

	expected := &dto.UploadFileResponse{
		Url:     "url",
		FileKey: "object",
	}

	t.Nil(err)
	t.Equal(expected, res)
}

func (t *FileServiceTest) TestUploadError() {
	id := uuid.New()
	objectKey := id.String()[:8] + "_" + "object"

	repo := mock_file.NewMockRepository(t.controller)
	repo.EXPECT().Upload([]byte("data"), "bucket", objectKey).Return("", "", errors.New("error"))

	svc := NewService(repo, t.logger)

	res, err := svc.Upload(&dto.UploadFileRequest{
		File: model.File{
			ID:   id,
			Name: "object",
			Data: []byte("data"),
		},
		Bucket: "bucket",
	})

	t.NotNil(err)
	t.Nil(res)
}

func (t *FileServiceTest) TestDeleteSuccess() {
	repo := mock_file.NewMockRepository(t.controller)
	repo.EXPECT().Delete("bucket", "object").Return(nil)

	svc := NewService(repo, t.logger)

	res, err := svc.Delete(&dto.DeleteFileRequest{
		FileKey: "object",
		Bucket:  "bucket",
	})

	expected := &dto.DeleteFileResponse{
		Success: true,
	}

	t.Nil(err)
	t.Equal(expected, res)
}

func (t *FileServiceTest) TestDeleteError() {
	repo := mock_file.NewMockRepository(t.controller)
	repo.EXPECT().Delete("bucket", "object").Return(errors.New("error"))

	svc := NewService(repo, t.logger)

	res, err := svc.Delete(&dto.DeleteFileRequest{
		FileKey: "object",
		Bucket:  "bucket",
	})

	t.NotNil(err)
	t.Nil(res)
}

func (t *FileServiceTest) TestGetSuccess() {
	repo := mock_file.NewMockRepository(t.controller)
	repo.EXPECT().Get("bucket", "object").Return("url", nil)

	svc := NewService(repo, t.logger)

	res, err := svc.Get(&dto.GetFileRequest{
		FileKey: "object",
		Bucket:  "bucket",
	})

	expected := &dto.GetFileResponse{
		FileUrl: "url",
	}

	t.Nil(err)
	t.Equal(expected, res)
}

func (t *FileServiceTest) TestGetError() {
	repo := mock_file.NewMockRepository(t.controller)
	repo.EXPECT().Get("bucket", "object").Return("", errors.New("error"))

	svc := NewService(repo, t.logger)

	res, err := svc.Get(&dto.GetFileRequest{
		FileKey: "object",
		Bucket:  "bucket",
	})

	t.NotNil(err)
	t.Nil(res)
}

func (t *FileServiceTest) TestGetNotFound() {
	repo := mock_file.NewMockRepository(t.controller)
	repo.EXPECT().Get("bucket", "object").Return("", nil)

	svc := NewService(repo, t.logger)

	res, err := svc.Get(&dto.GetFileRequest{
		FileKey: "object",
		Bucket:  "bucket",
	})

	t.NotNil(err)
	t.Nil(res)
}
