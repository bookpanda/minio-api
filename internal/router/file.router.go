package router

import "github.com/gin-gonic/gin"

func (r *Router) GetFile(path string, handler func(c *gin.Context)) {
	r.file.GET(path, handler)
}

func (r *Router) PostFile(path string, handler func(c *gin.Context)) {
	r.file.POST(path, handler)
}

func (r *Router) DeleteFile(path string, handler func(c *gin.Context)) {
	r.file.DELETE(path, handler)
}
