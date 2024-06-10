package router

import (
	"github.com/bookpanda/minio-api/config"
	"github.com/bookpanda/minio-api/internal/middleware"
	"github.com/gin-gonic/gin"
)

type Router struct {
	*gin.Engine
}

func New(conf *config.Config, corsHandler config.CorsHandler, appMiddleware middleware.AppMiddleware) *Router {
	if !conf.App.IsDevelopment() {
		gin.SetMode(gin.ReleaseMode)
	} else {
		gin.SetMode(gin.DebugMode)
	}

	r := gin.Default()
	r.Use(gin.HandlerFunc(corsHandler))
	r.Use(gin.HandlerFunc(appMiddleware))

	return &Router{r}
}
