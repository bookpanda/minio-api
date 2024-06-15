package middleware

import (
	"net/http"
	"strings"

	"github.com/bookpanda/minio-api/apperrors"
	"github.com/bookpanda/minio-api/config"
	"github.com/gin-gonic/gin"
)

type AppMiddleware gin.HandlerFunc

func NewAppMiddleware(conf *config.AppConfig) AppMiddleware {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, apperrors.Unauthorized)
			c.Abort()
			return
		}

		if !strings.HasPrefix(authHeader, "Bearer ") {
			c.JSON(http.StatusUnauthorized, apperrors.InvalidToken)
			c.Abort()
			return
		}

		token := strings.TrimPrefix(authHeader, "Bearer ")
		if token != conf.ApiKey {
			c.JSON(http.StatusUnauthorized, apperrors.InvalidToken)
			c.Abort()
			return
		}

		c.Next()
	}
}
