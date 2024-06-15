package file

import (
	"errors"
	"net/http"
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

	repo := NewRepository(t.conf, storeClient, nil)

	url, key, err := repo.Upload([]byte{}, "bucket", "object")
	t.Nil(err)
	t.Equal("object", key)
	t.Equal(repo.GetURL("bucket", "object"), url)
}

func (t *FileRepositoryTest) TestUploadError() {
	storeClient := store_client.NewMockClient(t.controller)
	storeClient.EXPECT().PutObject(gomock.Any(), "bucket", "object", gomock.Any(), int64(0), gomock.Any()).Return(minio.UploadInfo{}, errors.New("error"))

	repo := NewRepository(t.conf, storeClient, nil)

	url, key, err := repo.Upload([]byte{}, "bucket", "object")
	t.NotNil(err)
	t.Empty(url)
	t.Empty(key)
}

func (t *FileRepositoryTest) TestDeleteSuccess() {
	storeClient := store_client.NewMockClient(t.controller)
	storeClient.EXPECT().RemoveObject(gomock.Any(), "bucket", "object", gomock.Any()).Return(nil)

	repo := NewRepository(t.conf, storeClient, nil)

	err := repo.Delete("bucket", "object")
	t.Nil(err)
}

func (t *FileRepositoryTest) TestDeleteError() {
	storeClient := store_client.NewMockClient(t.controller)
	storeClient.EXPECT().RemoveObject(gomock.Any(), "bucket", "object", gomock.Any()).Return(errors.New("error"))

	repo := NewRepository(t.conf, storeClient, nil)

	err := repo.Delete("bucket", "object")
	t.NotNil(err)
}

func (t *FileRepositoryTest) TestGetSuccess() {
	httpClient := http_client.NewMockClient(t.controller)
	httpClient.EXPECT().Get("https://endpoint/bucket/object").Return(&http.Response{
		StatusCode: http.StatusOK,
	}, nil)

	repo := NewRepository(t.conf, nil, httpClient)

	url, err := repo.Get("bucket", "object")
	t.Nil(err)
	t.Equal(repo.GetURL("bucket", "object"), url)
}

func (t *FileRepositoryTest) TestGetError() {
	httpClient := http_client.NewMockClient(t.controller)
	httpClient.EXPECT().Get("https://endpoint/bucket/object").Return(nil, errors.New("error"))

	repo := NewRepository(t.conf, nil, httpClient)

	url, err := repo.Get("bucket", "object")
	t.NotNil(err)
	t.Empty(url)
}

func (t *FileRepositoryTest) TestGetNotFound() {
	httpClient := http_client.NewMockClient(t.controller)
	httpClient.EXPECT().Get("https://endpoint/bucket/object").Return(&http.Response{
		StatusCode: http.StatusNotFound,
	}, nil)

	repo := NewRepository(t.conf, nil, httpClient)

	url, err := repo.Get("bucket", "object")
	t.Nil(err)
	t.Empty(url)
}

func (t *FileRepositoryTest) TestGetURL() {
	repo := NewRepository(t.conf, nil, nil)
	url := repo.GetURL("bucket", "object")
	t.Equal("https://endpoint/bucket/object", url)
}
