package router

import (
	"github.com/bookpanda/minio-api/config"
	"github.com/bookpanda/minio-api/internal/middleware"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type Router struct {
	*gin.Engine
	file        *gin.RouterGroup
	healthCheck *gin.RouterGroup
}

func New(conf *config.Config, corsHandler config.CorsHandler, appMiddleware middleware.AppMiddleware) *Router {
	if !conf.App.IsDevelopment() {
		gin.SetMode(gin.ReleaseMode)
	} else {
		gin.SetMode(gin.DebugMode)
	}

	r := gin.Default()
	r.Use(gin.HandlerFunc(corsHandler))
	v1 := r.Group("/v1")

	if conf.App.IsDevelopment() {
		v1.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	}

	healthCheck := v1.Group("/hc")
	healthCheck.Use(gin.HandlerFunc(appMiddleware))

	file := v1.Group("/file")
	file.Use(gin.HandlerFunc(appMiddleware))

	return &Router{r, file, healthCheck}
}
