package qingflowapi

import (
	"fmt"
	"strconv"
	"time"
)

type DashApi struct {
	client Client
}

func NewDashApi(client Client) DashApi {
	return DashApi{client: client}
}

type DashAuth int

const (
	DASH_AUTH_WORKSPACE      DashAuth = 1
	DASH_AUTH_SPECIFIED_USER          = 2
	DASH_AUTH_PUBLIC                  = 3
)

type DashPublishStatus int

type Dash struct {
	DashKey           string            `json:"dashKey"`
	DashName          string            `json:"dashName"`
	DashIcon          string            `json:"dashIcon"`
	DashAuth          DashAuth          `json:"dashAuth"`
	AuthMembers       []AuthMember      `json:"authMembers"`
	Creator           CreatorUser       `json:"creator"`
	CreateTime        time.Time         `json:"createTime"`
	Tags              []AppTag          `json:"tags"`
	DashPublishStatus DashPublishStatus `json:"dashPublishStatus"`
}

type DashUpdationRequest struct {
	DashName    string       `json:"dashName"`
	TagIds      []ID         `json:"tagIds"`
	DashIcon    string       `json:"dashIcon"`
	DashAuth    DashAuth     `json:"dashAuth"`
	AuthMembers []AuthMember `json:"authMembers"`
}

type DashCreationReqeust struct {
	DashUpdationRequest
	UserId string `json:"userId"`
}

func (api DashApi) Create(request DashCreationReqeust) (string, error) {
	endpoint := "dash"
	var result ApiResponse[Dash]
	err := api.client.post(endpoint, request, &result)
	if err != nil {
		return "", err
	}
	return result.Result.DashKey, nil
}

func (api DashApi) Delete(dashKey string) (string, error) {
	endpoint := fmt.Sprintf("dash/%s", dashKey)
	var result ApiResponse[Dash]
	err := api.client.deleteRequest(endpoint, nil, &result)
	if err != nil {
		return "", err
	}
	return result.Result.DashKey, nil
}

func (api DashApi) Update(dashKey string, request DashUpdationRequest) (string, error) {
	endpoint := fmt.Sprintf("dash/%s", dashKey)
	var result ApiResponse[Dash]
	err := api.client.post(endpoint, request, &result)
	if err != nil {
		return "", err
	}
	return result.Result.DashKey, nil
}

func (api DashApi) Page(dashKey string, page int, size int) (PageResult[Dash], error) {
	endpoint := "dashes"
	params := map[string]string{
		"dashKey":  dashKey,
		"pageNum":  strconv.Itoa(page),
		"pageSize": strconv.Itoa(size),
	}
	var result ApiResponse[PageResult[Dash]]
	err := api.client.get(endpoint, params, &result)
	if err != nil {
		return PageResult[Dash]{}, err
	}
	return result.Result, nil
}

func (api DashApi) GetAll() ([]Dash, error) {
	endpoint := "dash"
	var result ApiResponse[struct {
		DashList []Dash `json:"dashList"`
	}]
	err := api.client.get(endpoint, nil, &result)
	if err != nil {
		return []Dash{}, err
	}
	return result.Result.DashList, nil
}
