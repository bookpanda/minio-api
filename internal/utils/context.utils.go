package utils

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"mime/multipart"
	"strings"
)

func ExtractFile(file *multipart.FileHeader, allowedContent map[string]struct{}, maxSize int64) (data []byte, err error) {
	if !isExisted(allowedContent, file.Header["Content-Type"][0]) {
		return nil, errors.New("Allowed content type is " + fmt.Sprint(strings.Join(mapToArr(allowedContent), ", ")))
	}

	if file.Size > maxSize*1000000000 {
		return nil, fmt.Errorf("Max file size is %v", maxSize)
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
	return ok
}

func mapToArr(m map[string]struct{}) []string {
	var arr []string
	for k := range m {
		arr = append(arr, k)
	}
	return arr
}
