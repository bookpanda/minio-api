package cmd

import (
	"fmt"

	"github.com/bookpanda/minio-api/config"
	"github.com/bookpanda/minio-api/internal/file"
	healthcheck "github.com/bookpanda/minio-api/internal/health_check"
	"github.com/bookpanda/minio-api/internal/middleware"
	"github.com/bookpanda/minio-api/internal/router"
	"github.com/bookpanda/minio-api/logger"
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

	logger := logger.New(conf)
	corsHandler := config.MakeCorsConfig(conf)
	appMiddleware := middleware.NewAppMiddleware(&conf.App)

	hcHandler := healthcheck.NewHandler()

	fileRepo := file.NewRepository(conf.Store, minioClient)
	fileSvc := file.NewService(fileRepo, logger)
	fileHdr := file.NewHandler(fileSvc, logger)

	r := router.New(conf, corsHandler, appMiddleware)

	r.GET("/hc", hcHandler.HealthCheck)
	r.GET("/get", fileHdr.Get)
	r.POST("/upload", fileHdr.Upload)
	r.DELETE("/delete", fileHdr.Delete)
}
