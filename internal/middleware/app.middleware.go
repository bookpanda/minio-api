package middleware

import (
	"strings"

	"github.com/bookpanda/minio-api/config"
	"github.com/bookpanda/minio-api/errors"
	"github.com/gin-gonic/gin"
)

type AppMiddleware gin.HandlerFunc

func NewAppMiddleware(conf *config.AppConfig) AppMiddleware {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			errors.ResponseError(c, errors.Unauthorized)
			c.Abort()
			return
		}

		if !strings.HasPrefix(authHeader, "Bearer ") {
			errors.ResponseError(c, errors.InvalidToken)
			c.Abort()
			return
		}

		token := strings.TrimPrefix(authHeader, "Bearer ")
		if token != conf.ApiKey {
			errors.ResponseError(c, errors.InvalidToken)
			c.Abort()
			return
		}

		c.Next()
	}
}
