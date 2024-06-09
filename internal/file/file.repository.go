package file

import (
	"github.com/bookpanda/minio-api/config"
	"github.com/minio/minio-go/v7"
)

type Repository interface {
	Upload()
	Delete()
	Get()
}

type repositoryImpl struct {
	conf  config.StoreConfig
	minio *minio.Client
}

func NewRepository(conf config.StoreConfig, minioClient *minio.Client) Repository {
	return &repositoryImpl{
		conf:  conf,
		minio: minioClient,
	}
}

func (r *repositoryImpl) Upload() {
	// r.minio.
}

func (r *repositoryImpl) Delete() {}

func (r *repositoryImpl) Get() {}
