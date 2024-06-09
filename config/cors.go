package config

import (
	"strings"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

type CorsHandler gin.HandlerFunc

func makeCorsConfig(config *Config) gin.HandlerFunc {
	if config.App.IsDevelopment() {
		return cors.New(cors.Config{
			AllowMethods:     []string{"*"},
			AllowHeaders:     []string{"*"},
			AllowCredentials: true,
			MaxAge:           12 * time.Hour,
			AllowOriginFunc: func(string) bool {
				return true
			},
		})

	}

	allowOrigins := strings.Split(config.Cors.AllowOrigins, ",")

	return cors.New(cors.Config{
		AllowOrigins:     allowOrigins,
		AllowMethods:     []string{"*"},
		AllowHeaders:     []string{"*"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	})
}

func MakeCorsConfig(cfg *Config) CorsHandler {
	return CorsHandler(makeCorsConfig(cfg))
}
