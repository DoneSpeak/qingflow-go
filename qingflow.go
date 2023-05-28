package qingflow

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

// https://qingflow.com/help/docs/6239a267c7c55500418da348
const (
	QUE_TYPE_DESC       int = 1
	QUE_TYPE_LINE           = 2
	QUE_TYPE_MULTI_LINE     = 3
	QUE_TYPE_DATE           = 4
	QUE_TYPE_MEMBER         = 5
)

// https://qingflow.com/help/docs/6239a267c7c55500418da348
const (
	JUDGE_TYPE_NE     int = 1
	JUDGE_TYPE_IN         = 2
	JUDGE_TYPE_NOT_IN     = 3
)

type ApiResponse[R any] struct {
	ErrCode           int      `json:"errCode"`
	ErrMsg            string   `json:"errMsg"`
	QuestionRelations []string `json:"questionRelations"`
	Result            R        `json:"result"`
}

type PageResult[R any] struct {
	PageAmount int `json:"pageAmount"`
	PageNum    int `json:"pageNum"`
	PageSize   int `json:"pageSize"`
	Result     []R `json:"result"`
}

type Client struct {
	BaseUrl    string
	Token      AccessToken
	HttpClient http.Client
}

func (c Client) createUrl(path string) string {
	// TODO 考虑多了斜杠的情况
	return c.BaseUrl + "/" + path
}

func (c Client) get(path string, result any) error {
	url := c.createUrl(path)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return err
	}
	err = c.executeRequest(req, result)
	return err
}

func (c Client) post(path string, body any, result any) error {
	return c.execute("POST", path, body, result)
}

func (c Client) put(path string, body any, result any) error {
	return c.execute("PUT", path, body, result)
}

func (c Client) delete(path string) error {
	var result ApiResponse[any]
	err := c.execute("DELETE", path, nil, result)
	if err != nil {
		return err
	}
	fmt.Println("response: ", result)
	return err
}

func (c Client) execute(method string, path string, body any, result any) error {
	var jsonBytes io.Reader
	if body != nil {
		jsonBody, err := json.Marshal(body)
		if err != nil {
			return err
		}
		fmt.Println("Requets: ", string(jsonBody))
		jsonBytes = bytes.NewBuffer(jsonBody)
	}
	url := c.createUrl(path)
	fmt.Println("Request URL: ", url)
	request, err := http.NewRequest(method, url, jsonBytes)
	if err != nil {
		return err
	}
	return c.executeRequest(request, result)
}

func (c Client) executeRequest(request *http.Request, result any) error {
	request.Header.Add("Accept", `application/json`)
	request.Header.Add("Content-Type", `application/json`)
	request.Header.Add("accessToken", c.Token.getValue())
	resp, err := c.HttpClient.Do(request)
	if err != nil {
		// transfer to customized error
		return err
	}
	defer resp.Body.Close()
	body, _ := io.ReadAll(resp.Body)
	fmt.Println("Body: ", string(body))
	json.Unmarshal(body, result)
	return nil
}

func defaultClient(baseUrl string) Client {
	HttpClient := http.Client{Timeout: time.Duration(30) * time.Second}
	return Client{BaseUrl: baseUrl, HttpClient: HttpClient}
}

func (c Client) User() UserApi {
	return UserApi{client: c}
}

func (c Client) Apply(appKey string) ApplyApi {
	return ApplyApi{client: c, appKey: appKey}
}
