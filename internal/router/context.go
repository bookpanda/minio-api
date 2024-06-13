package router

import (
	"github.com/bookpanda/minio-api/errors"
	"github.com/bookpanda/minio-api/internal/dto"
	"github.com/bookpanda/minio-api/internal/utils"
	"github.com/gin-gonic/gin"
)

type Context interface {
	JSON(statusCode int, obj interface{})
	ResponseError(err *errors.AppError)
	Bind(obj interface{}) error
	PostForm(key string) string
	FormFile(key string, allowedContentType map[string]struct{}, maxFileSize int64) (*dto.DecomposedFile, error)
}

type contextImpl struct {
	*gin.Context
}

func NewContext(c *gin.Context) Context {
	return &contextImpl{c}
}

func (c *contextImpl) JSON(statusCode int, obj interface{}) {
	c.Context.JSON(statusCode, obj)
}

func (c *contextImpl) ResponseError(err *errors.AppError) {
	c.JSON(err.HttpCode, gin.H{"error": err.Error()})
}

func (c *contextImpl) Bind(obj interface{}) error {
	return c.Context.Bind(obj)
}

func (c *contextImpl) PostForm(key string) string {
	return c.Context.PostForm(key)
}

func (c *contextImpl) FormFile(key string, allowedContentType map[string]struct{}, maxFileSize int64) (*dto.DecomposedFile, error) {
	file, err := c.Context.FormFile(key)
	if err != nil {
		return nil, err
	}

	data, err := utils.ExtractFile(file, allowedContentType, maxFileSize)
	if err != nil {
		return nil, err
	}

	return &dto.DecomposedFile{
		Filename: file.Filename,
		Data:     data,
	}, nil
}
