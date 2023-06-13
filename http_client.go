package qingflowapi

import (
	"net/http"
)

type HttpClient interface {
	execute(request *http.Request) (*http.Response, error)
}

type NetHttpClient struct {
	httpClient http.Client
}

func (c NetHttpClient) execute(request *http.Request) (*http.Response, error) {
	request.Header.Add("Accept", `application/json`)
	request.Header.Add("Content-Type", `application/json`)
	return c.httpClient.Do(request)
}

type ForwardingHttpClient struct {
	httpClient HttpClient
}

func (c ForwardingHttpClient) execute(request *http.Request) (*http.Response, error) {
	return c.httpClient.execute(request)
}
