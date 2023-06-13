package qingflowapi

import (
	"fmt"
	"time"
)

type AccessToken interface {
	getValue() string
	getExpireAt() time.Time
}

type SimpleAccessToken struct {
	AccessToken string `json:"accessToken"`
	ExpireTime  int
	ExpireAt    time.Time `json:"expireAt"`
}

func (t SimpleAccessToken) getValue() string {
	return t.AccessToken
}

func (t SimpleAccessToken) getExpireAt() time.Time {
	return t.ExpireAt
}

type AutoRefreshAccessToken struct {
	authApi AuthApi
	cred    Credential
	token   AccessToken
}

func (t AutoRefreshAccessToken) getValue() (string, error) {
	if t.token.getExpireAt().Before(time.Now()) {
		token, err := t.authApi.GrantToken(t.cred)
		if err != nil {
			return "", err
		}
		t.token = token
	}
	return t.token.getValue(), nil
}

type AuthApi struct {
	accessToken AccessToken
	client      Client
}

type Credential struct {
	WsId     string
	WsSecret string
}

func (api AuthApi) GrantToken(cred Credential) (AccessToken, error) {
	path := fmt.Sprintf("/accessToken?wsId=%s&wsSecret=%s", cred.WsId, cred.WsSecret)
	var resp ApiResponse[SimpleAccessToken]
	err := api.client.get(path, &resp)
	if err != nil {
		return nil, err
	}
	token := resp.Result
	token.ExpireAt = time.Now().Local().Add(time.Second * time.Duration(token.ExpireTime))

	return token, nil
}
