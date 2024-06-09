package cmd

import (
	"fmt"

	"github.com/bookpanda/minio-api/config"
	"github.com/bookpanda/minio-api/internal/file"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

func main() {
	conf, err := config.LoadConfig()
	if err != nil {
		panic(fmt.Sprintf("Failed to load config: %v", err))
	}

	minioClient, err := minio.New(conf.Store.Endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(conf.Store.AccessKey, conf.Store.SecretKey, ""),
		Secure: conf.Store.UseSSL,
	})

	fileRepo := file.NewRepository(conf.Store, minioClient)
	fileSvc := file.NewService(fileRepo)
}
