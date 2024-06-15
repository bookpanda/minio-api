package file

import (
	"testing"

	"github.com/bookpanda/minio-api/internal/dto"
	"github.com/bookpanda/minio-api/internal/model"
	mock_file "github.com/bookpanda/minio-api/mocks/repository"
	"github.com/golang/mock/gomock"
	"github.com/google/uuid"
	"github.com/stretchr/testify/suite"
)

type FileServiceTest struct {
	suite.Suite
	controller *gomock.Controller
}

func TestFileService(t *testing.T) {
	suite.Run(t, new(FileServiceTest))
}

func (t *FileServiceTest) SetupTest() {
	t.controller = gomock.NewController(t.T())
}

func (t *FileServiceTest) TestUploadSuccess() {
	id := uuid.New()
	objectKey := id.String()[:8] + "_" + "object"

	repo := mock_file.NewMockRepository(t.controller)
	repo.EXPECT().Upload([]byte("data"), "bucket", objectKey).Return("url", "object", nil)

	svc := NewService(repo, nil)

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
