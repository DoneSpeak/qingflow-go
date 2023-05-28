package qingflow

import (
	"net/http"
	"time"
)

type AccessToken interface {
	getValue() string
	getExpireAt() time.Time
}

type SimpleAccessToken struct {
	Value    string
	ExpireAt time.Time
}

func (t SimpleAccessToken) getValue() string {
	return t.Value
}

func (t SimpleAccessToken) getExpireAt() time.Time {
	return t.ExpireAt
}

type AuthHttpClient struct {
	accessToken AccessToken
	httpClient  HttpClient
}

func (c AuthHttpClient) execute(request *http.Request) (*http.Response, error) {
	request.Header.Add("accessToken", c.accessToken.getValue())
	return c.httpClient.execute(request)
}
