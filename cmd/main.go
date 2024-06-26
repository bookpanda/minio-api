package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/bookpanda/minio-api/config"
	"github.com/bookpanda/minio-api/constants"
	store_client "github.com/bookpanda/minio-api/internal/client/store"
	fileHdr "github.com/bookpanda/minio-api/internal/handler/file"
	healthcheck "github.com/bookpanda/minio-api/internal/health_check"
	"github.com/bookpanda/minio-api/internal/middleware"
	fileRepo "github.com/bookpanda/minio-api/internal/repository/file"
	"github.com/bookpanda/minio-api/internal/router"
	fileSvc "github.com/bookpanda/minio-api/internal/service/file"
	"github.com/bookpanda/minio-api/internal/validator"
	"github.com/bookpanda/minio-api/logger"
	"github.com/bookpanda/minio-api/metrics"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

// @title           Minio API
// @version         1.0
// @description     Object store API for personal projects

// @host      localhost:3000
// @BasePath  /api/v1
// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
// @description Type "Bearer" followed by a space and JWT token.
func main() {
	conf, err := config.LoadConfig()
	if err != nil {
		panic(fmt.Sprintf("Failed to load config: %v", err))
	}

	minioClient, err := minio.New(conf.Store.Endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(conf.Store.AccessKey, conf.Store.SecretKey, ""),
		Secure: conf.Store.UseSSL,
	})
	if err != nil {
		panic(fmt.Sprintf("Failed to connect to Minio: %v", err))
	}
	storeClient := store_client.NewClient(minioClient)

	httpClient := &http.Client{}

	validator, err := validator.NewDtoValidator()
	if err != nil {
		panic(fmt.Sprintf("Failed to create dto validator: %v", err))
	}

	logger := logger.New(conf)
	corsHandler := config.MakeCorsConfig(conf)
	appMiddleware := middleware.NewAppMiddleware(&conf.App)

	requestsCounter := prometheus.NewCounterVec(prometheus.CounterOpts{
		Name: "api_requests_total",
		Help: "Total number of API requests by domain, method and status code",
	}, []string{"domain", "method", "status_code"})
	requestsMetrics := metrics.NewRequestsMetrics(requestsCounter)

	metricsRegistry := prometheus.NewRegistry()
	metrics := metrics.NewMetrics(metricsRegistry, requestsMetrics)

	hcHandler := healthcheck.NewHandler()

	fileRepo := fileRepo.NewRepository(&conf.Store, storeClient, httpClient)
	fileSvc := fileSvc.NewService(fileRepo, logger)
	fileHdr := fileHdr.NewHandler(fileSvc, validator, conf.App.MaxFileSizeMB, constants.AllowedContentType, logger, requestsMetrics)

	r := router.New(conf, corsHandler, appMiddleware)

	r.GetHealthCheck("/", hcHandler.HealthCheck)
	r.GetMetrics("/", promhttp.HandlerFor(
		metrics.Registry(),
		promhttp.HandlerOpts{
			EnableOpenMetrics: true,
		}),
	)
	r.PostFile("/upload", fileHdr.Upload)
	r.GetFile("/get/:bucket", fileHdr.Get)
	r.DeleteFile("/delete/:bucket", fileHdr.Delete)

	srv := &http.Server{
		Addr:    fmt.Sprintf(":%v", conf.App.Port),
		Handler: r.Handler(),
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutdown Server ...")

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown:", err)
	}
	<-ctx.Done()
	log.Println("timeout of 3 seconds.")

	log.Println("Server exiting")
}
