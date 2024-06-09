package middleware

import (
	"strings"

	"github.com/bookpanda/minio-api/config"
	"github.com/bookpanda/minio-api/errors"
	"github.com/gin-gonic/gin"
)

type AppMidddleware gin.HandlerFunc

func NewAppMiddleware(conf *config.Config) AppMidddleware {
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
		if token != conf.App.ApiKey {
			errors.ResponseError(c, errors.InvalidToken)
			c.Abort()
			return
		}

		c.Next()
	}
}
