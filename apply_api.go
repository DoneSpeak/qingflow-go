package qingflow

import (
	"fmt"
)

type ApplyApi struct {
	client Client
	appKey string
}

func (api ApplyApi) Query(query ApplyQuery) (PageResult[Apply], error) {
	path := fmt.Sprintf("app/%s/apply/filter", api.appKey)
	var result ApiResponse[PageResult[Apply]]
	err := api.client.post(path, query, &result)
	if err != nil {
		return PageResult[Apply]{}, err
	}
	return result.Result, nil
}

func (api ApplyApi) Create(request CreateRequest) (CreateResponse, error) {
	path := fmt.Sprintf("app/%s/apply", api.appKey)
	var result ApiResponse[CreateResponse]
	err := api.client.post(path, request, &result)
	if err != nil {
		return CreateResponse{}, err
	}
	return result.Result, nil
}

func (api ApplyApi) FindById(applyId string) (Apply, error) {
	path := fmt.Sprintf("app/%s/apply/%s", api.appKey, applyId)
	var result ApiResponse[Apply]
	api.client.get(path, &result)
	return result.Result, nil
}

type ApplyQuery struct {
	PageSize   int `json:"pageSize"`
	PageNumber int `json:"pageNumber"`
	Type       int `json:"type"`
	Sorts      []struct {
		QueId    int  `json:"queId"`
		IsAscend bool `json:"isAscend"`
	} `json:"sorts"`
	Queries []struct {
		QueId         int      `json:"queId"`
		SearchKey     string   `json:"searchKey"`
		SearchKeys    []string `json:"searchKeys"`
		MinValue      string   `json:"minValue"`
		MaxValue      string   `json:"maxValue"`
		Scope         int      `json:"scope"`
		SearchOptions []int    `json:"searchOptions"`
		SearchUserIds []string `json:"searchUserIds"`
	} `json:"queries"`
	QueryKey string `json:"queryKey"`
	ApplyIds []int  `json:"applyIds"`
}

type Apply struct {
	ApplyId       string   `json:"applyId"`
	Answers       []Answer `json:"answers"`
	ApplyBaseInfo string   `json:"applyBaseInfo"`
}

type Answer struct {
	QueId       int      `json:"queId"`
	QueTitle    string   `json:"queTitle"`
	QueType     int      `json:"queType"`
	TableValues []string `json:"tableValues"`
	Values      []struct {
		Value     string `json:"value"`
		OtherInfo string `json:"otherInfo"`
		Id        string `json:"id"`
	} `json:"values"`
}

type CreateRequest struct{}

type CreateResponse struct{}
