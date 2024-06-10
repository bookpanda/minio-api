package file

import (
	"bytes"
	"context"
	"fmt"
	"time"

	"github.com/bookpanda/minio-api/config"
	"github.com/minio/minio-go/v7"
	"github.com/pkg/errors"
)

type Repository interface {
	Upload(file []byte, bucketName string, objectKey string) (url string, key string, err error)
	Delete(bucketName string, objectKey string) (err error)
	Get(bucketName string, objectKey string) (url string, err error)
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

func (r *repositoryImpl) Upload(file []byte, bucketName string, objectKey string) (url string, key string, err error) {
	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, 50*time.Second)
	defer cancel()

	buffer := bytes.NewReader(file)

	uploadOutput, err := r.minio.PutObject(context.Background(), bucketName, objectKey, buffer,
		buffer.Size(), minio.PutObjectOptions{ContentType: "application/octet-stream"})
	if err != nil {
		return "", "", errors.Wrap(err, fmt.Sprintf("Couldn't upload object to %v/%v.", bucketName, objectKey))
	}

	return r.getURL(bucketName, objectKey), uploadOutput.Key, nil
}

func (r *repositoryImpl) Delete(bucketName string, objectKey string) (err error) {
	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, 50*time.Second)
	defer cancel()

	opts := minio.RemoveObjectOptions{
		GovernanceBypass: true,
	}
	err = r.minio.RemoveObject(context.Background(), bucketName, objectKey, opts)
	if err != nil {
		return errors.Wrap(err, fmt.Sprintf("Couldn't delete object %v/%v.", bucketName, objectKey))
	}

	return nil
}

func (r *repositoryImpl) Get(bucketName string, objectKey string) (url string, err error) {
	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, 50*time.Second)
	defer cancel()

	_, err = r.minio.StatObject(ctx, bucketName, objectKey, minio.StatObjectOptions{})
	if err != nil {
		return "", errors.Wrap(err, fmt.Sprintf("Couldn't get object %v/%v.", bucketName, objectKey))
	}

	return r.getURL(bucketName, objectKey), nil
}

func (c *repositoryImpl) getURL(bucketName string, objectKey string) string {
	return "https://" + c.conf.Endpoint + "/" + bucketName + "/" + objectKey
}
