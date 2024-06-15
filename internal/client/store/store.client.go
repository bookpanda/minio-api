package store

import (
	"context"
	"io"

	"github.com/minio/minio-go/v7"
)

type Client interface {
	PutObject(ctx context.Context, bucketName string, objectName string, reader io.Reader, objectSize int64, opts minio.PutObjectOptions) (info minio.UploadInfo, err error)
	RemoveObject(ctx context.Context, bucketName string, objectName string, opts minio.RemoveObjectOptions) error
}

type clientImpl struct {
	*minio.Client
}

func NewClient(minioClient *minio.Client) Client {
	return &clientImpl{minioClient}
}

func (c *clientImpl) PutObject(ctx context.Context, bucketName string, objectName string, reader io.Reader, objectSize int64, opts minio.PutObjectOptions) (info minio.UploadInfo, err error) {
	return c.Client.PutObject(ctx, bucketName, objectName, reader, objectSize, opts)
}

func (c *clientImpl) RemoveObject(ctx context.Context, bucketName string, objectName string, opts minio.RemoveObjectOptions) error {
	return c.Client.RemoveObject(ctx, bucketName, objectName, opts)
}
