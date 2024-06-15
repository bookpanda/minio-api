package file

import (
	"errors"
	"testing"

	"github.com/bookpanda/minio-api/config"
	http_client "github.com/bookpanda/minio-api/mocks/client/http"
	store_client "github.com/bookpanda/minio-api/mocks/client/store"
	"github.com/golang/mock/gomock"
	"github.com/minio/minio-go/v7"
	"github.com/stretchr/testify/suite"
)

type FileRepositoryTest struct {
	suite.Suite
	conf       *config.StoreConfig
	controller *gomock.Controller
}

func TestFileRepository(t *testing.T) {
	suite.Run(t, new(FileRepositoryTest))
}

func (t *FileRepositoryTest) SetupTest() {
	t.conf = &config.StoreConfig{
		Endpoint: "endpoint",
	}
	t.controller = gomock.NewController(t.T())

}

func (t *FileRepositoryTest) TestUploadSuccess() {
	storeClient := store_client.NewMockClient(t.controller)
	storeClient.EXPECT().PutObject(gomock.Any(), "bucket", "object", gomock.Any(), int64(0), gomock.Any()).Return(minio.UploadInfo{Key: "object"}, nil)

	httpClient := http_client.NewMockClient(t.controller)

	repo := NewRepository(t.conf, storeClient, httpClient)

	url, key, err := repo.Upload([]byte{}, "bucket", "object")
	t.Nil(err)
	t.Equal("object", key)
	t.Equal("https://endpoint/bucket/object", url)
}

func (t *FileRepositoryTest) TestUploadError() {
	storeClient := store_client.NewMockClient(t.controller)
	storeClient.EXPECT().PutObject(gomock.Any(), "bucket", "object", gomock.Any(), int64(0), gomock.Any()).Return(minio.UploadInfo{}, errors.New("error"))

	httpClient := http_client.NewMockClient(t.controller)

	repo := NewRepository(t.conf, storeClient, httpClient)

	url, key, err := repo.Upload([]byte{}, "bucket", "object")
	t.NotNil(err)
	t.Empty(url)
	t.Empty(key)
}
