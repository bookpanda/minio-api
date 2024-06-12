package utils

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"mime/multipart"
)

func ExtractFile(file *multipart.FileHeader, allowContent map[string]struct{}, maxSize int64) (data []byte, err error) {
	if !isExisted(allowContent, file.Header["Content-Type"][0]) {
		return nil, errors.New("Not allowed content")
	}

	if file.Size > maxSize {
		return nil, errors.New(fmt.Sprintf("Max file size is %v", maxSize))
	}

	fileBytes, err := file.Open()
	if err != nil {
		return nil, err
	}

	defer fileBytes.Close()

	buf := bytes.NewBuffer(nil)
	if _, err := io.Copy(buf, fileBytes); err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

func isExisted(e map[string]struct{}, key string) bool {
	_, ok := e[key]
	if ok {
		return true
	}
	return false
}
