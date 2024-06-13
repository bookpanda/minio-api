package router

import "github.com/gin-gonic/gin"

func (r *Router) GetFile(path string, handler func(c Context)) {
	r.file.GET(path, func(c *gin.Context) {
		handler(NewContext(c))
	})
}

func (r *Router) PostFile(path string, handler func(c Context)) {
	r.file.POST(path, func(c *gin.Context) {
		handler(NewContext(c))
	})
}

func (r *Router) DeleteFile(path string, handler func(c Context)) {
	r.file.DELETE(path, func(c *gin.Context) {
		handler(NewContext(c))
	})
}
