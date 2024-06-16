package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (r *Router) GetMetrics(path string, handler http.Handler) {
	r.metrics.GET(path, func(c *gin.Context) {
		handler.ServeHTTP(c.Writer, c.Request)
	})
}
