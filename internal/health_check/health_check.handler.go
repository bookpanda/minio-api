package healthcheck

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Handler interface {
	HealthCheck(c *gin.Context)
}

func NewHandler() Handler {
	return &handlerImpl{}
}

type handlerImpl struct {
}

func (h *handlerImpl) HealthCheck(c *gin.Context) {
	c.String(http.StatusOK, "OK")
}
