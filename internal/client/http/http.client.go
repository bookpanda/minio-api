package http

import (
	"net/http"
)

type Client interface {
	Get(url string) (resp *http.Response, err error)
}

type clientImpl struct {
	*http.Client
}

func NewClient(httpClient *http.Client) Client {
	return &clientImpl{httpClient}
}

func (c *clientImpl) Get(url string) (resp *http.Response, err error) {
	return c.Client.Get(url)
}
