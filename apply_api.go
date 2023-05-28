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
	PageSize   int
	PageNumber int
	Type       int
	Sorts      []struct {
		QueId    int
		IsAscend bool
	}
	Queries []struct {
		QueId         int
		SearchKey     string
		SearchKeys    []string
		MinValue      string
		MaxValue      string
		Scope         int
		SearchOptions []int
		SearchUserIds []string
	}
	QueryKey string
	ApplyIds []int
}

type Apply struct{}

type Answer struct {
	AssociateQueType int
}

type CreateRequest struct{}

type CreateResponse struct{}
