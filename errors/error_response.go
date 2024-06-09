package errors

import "github.com/gin-gonic/gin"

func ResponseError(c *gin.Context, err *AppError) {
	c.JSON(err.HttpCode, gin.H{"error": err.Error()})
}
