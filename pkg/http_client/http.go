package http_client

import (
	"net/http"
)

const UserAgent = "health-check"

type HttpClient interface {
	SimpleGet(url string, body map[string]string, res interface{}) error
	SimplePost(url string, body interface{}, res interface{}) error
	HttpGet(url string, body map[string]string) ([]byte, http.Header, int, error)
	HttpPost(url string, body interface{}, headers map[string]string) ([]byte, http.Header, int, error)
}

type httpClient struct {
}

func NewHttpClient() HttpClient {
	return &httpClient{}
}
