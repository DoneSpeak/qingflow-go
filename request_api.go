package qingflowapi

import "fmt"

type RequestApi struct {
	client Client
}

func NewRequestApi(client Client) RequestApi {
	return RequestApi{client: client}
}

type RequestResult struct {
	SuccessUsers []struct {
		UserId           SID
		Email            string
		CustomRole       []ID // 自建角色id列表（仅企业微信/钉钉/飞书等特殊工作区适用）
		CustomDepartment []ID // 自建部门id列表（仅企业微信/钉钉/飞书等特殊工作区适用）
		BeingDisabled    bool // 是否禁用成员
		BeingActive      bool // 用户的激活状态
	}
	ErrorUsers []struct {
		ErrorCode        string
		UserId           SID
		Email            string
		CustomRole       []ID
		CustomDepartment []ID
	}
	MemberInfo []any
}

func (api RequestApi) Get(requestId RequestID) (RequestResult, error) {
	endpoint := fmt.Sprintf("operation/%s", requestId)
	var result ApiResponse[RequestResult]
	err := api.client.get(endpoint, nil, &result)
	if err != nil {
		return RequestResult{}, err
	}
	return result.Result, nil
}
