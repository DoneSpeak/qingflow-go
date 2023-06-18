package qingflowapi

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

type ID int64

type SID string

type RequestID string

type QuestionRelation struct {
	QueId             string      `json:"queId"`
	MatchRuleType     int         `json:"matchRuleType"`
	MatchRules        []MatchRule `json:"matchRules"`
	MatchRuleFormulat string      `json:"matchRuleFormulat"`
	TableMatchRules   []struct {
		SubQueId        int `json:"subQueId"`
		RelatedSubQueId int `json:"relatedSubQueId"`
	} `json:"tableMatchRules"`
}

type MatchRule struct {
	QueId            int      `json:"queId"`
	QueTitle         string   `json:"queTitle"`
	JudgeType        int      `json:"judgeType"`
	MatchType        int      `json:"matchType"`
	JudgeValueQueIds []string `json:"judgeValueQueIds"`
}

type AnswerValue struct {
	// 问题的结果值
	Value string
	// 问题结果的其他信息,上传文件这里是文件名,选择类的其他选项这里是其他选项填写的内容
	OtherInfo string
	// 选择类型，为optId，成员类型，为uid；成员字段 ，当前用户: -1；地址字段，1-4：省、市、区、详细地址
	Id int
}

type AuthMember struct {
	UserIds           []SID `json:"userIds"`           // 可见的用户列表；外部用户id列表
	DeptIds           []ID  `json:"deptIds"`           // 可见的部门列表；外部部门id列表
	RoleIds           []ID  `json:"roleIds"`           // 可见的角色列表；外部角色id列表
	IncludeSubDeparts bool  `json:"includeSubDeparts"` // (应用已经使用相同字段)是否动态包含已选部门下的子部门，默认包含
}

type CreatorUser struct {
	UserId   string `json:"userId"`
	NickName string `json:"nickName"`
	HeadImg  string `json:"headImg"`
}

type ApiErrorResponse struct {
	ErrCode int    `json:"errCode"`
	ErrMsg  string `json:"errMsg"`
}

type ApiResponse[R any] struct {
	QuestionRelations []string `json:"questionRelations"`
	Result            R        `json:"result"`
}

type PageResult[R any] struct {
	PageAmount   int `json:"pageAmount"`
	PageNum      int `json:"pageNum"`
	PageSize     int `json:"pageSize"`
	ResultAmount int `json:"resultAmount"`
	Result       []R `json:"result"`
}

type Client struct {
	baseUrl    *url.URL
	token      AccessToken
	userId     string
	httpClient *http.Client
}

type ClientOption func(*Client)

func NewClient(baseUrl string, token AccessToken, opts ...ClientOption) Client {
	u, err := url.Parse(baseUrl)
	if err != nil {
		panic(err)
	}
	client := Client{
		baseUrl:    u,
		token:      token,
		httpClient: http.DefaultClient,
	}
	for _, opt := range opts {
		opt(&client)
	}
	return client
}

func (c Client) createUrl(path string) *url.URL {
	result, err := c.baseUrl.Parse(path)
	if err != nil {
		panic(err)
	}
	return result
}

func (c Client) get(path string, params map[string]string, result any) error {
	return c.execute("GET", path, nil, nil, result)
}

func (c Client) post(path string, body any, result any) error {
	return c.execute("POST", path, nil, body, result)
}

func (c Client) put(path string, body any, result any) error {
	return c.execute("PUT", path, nil, body, result)
}

func (c Client) delete(path string) error {
	var result *ApiResponse[any]
	return c.execute("DELETE", path, nil, nil, result)
}

func (c Client) deleteRequest(path string, body any, result any) error {
	return c.execute("DELETE", path, nil, body, result)
}

func (c Client) execute(method string, path string, params map[string]string, body any, result any) error {
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
	if params != nil && len(params) > 0 {
		q := url.Query()
		for k, v := range params {
			q.Add(k, v)
		}
		url.RawQuery = q.Encode()
	}
	request, err := http.NewRequest(method, url.String(), jsonBytes)
	fmt.Println("Request URL: ", url)
	if err != nil {
		return err
	}
	return c.executeRequest(request, result)
}

func (c Client) executeRequest(request *http.Request, result any) error {
	request.Header.Add("Accept", `application/json`)
	request.Header.Add("Content-Type", `application/json`)
	if c.token != nil {
		request.Header.Add("accessToken", c.token.getValue())
	}
	if len(c.userId) > 0 {
		request.Header.Add("userId", c.userId)
	}
	resp, err := c.httpClient.Do(request)
	if err != nil {
		// transfer to customized error
		return err
	}
	defer resp.Body.Close()
	body, _ := io.ReadAll(resp.Body)
	fmt.Println(string(body[:]))
	var apiErrResp ApiErrorResponse
	err = json.Unmarshal(body, &apiErrResp)
	if err != nil {
		return err
	}
	err = translateError(apiErrResp.ErrCode, apiErrResp.ErrMsg)
	if err != nil {
		return err
	}
	return json.Unmarshal(body, result)
}

func DefaultClient(token AccessToken) Client {
	baseUrl := "https://api.qingflow.com"
	return NewClient(baseUrl, token)
}

func (c Client) User() UserApi {
	return UserApi{client: c}
}

func (c Client) Apply(appKey string) ApplyApi {
	return ApplyApi{client: c, appKey: appKey}
}

func (c Client) Auth() AuthApi {
	return AuthApi{client: c}
}
