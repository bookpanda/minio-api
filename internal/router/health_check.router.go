package router

import "github.com/gin-gonic/gin"

func (r *Router) GetHealthCheck(path string, handler func(c *gin.Context)) {
	r.healthCheck.GET(path, handler)
}
