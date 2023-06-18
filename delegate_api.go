package qingflowapi

import (
	"strconv"
	"time"
)

type DelegateApi struct {
	client Client
}

func NewDelegateApi(client Client) DelegateApi {
	return DelegateApi{client: client}
}

type DelegateStatus int

const (
	DELEGATE_STATUS_NOT_START   DelegateStatus = 1
	DELEGATE_STATUS_DELEGATIHNG                = 2
	DELEGATE_STATUS_DONE                       = 3
)

type Delegate struct {
	Id                 ID        `json:"id"`
	CreateTime         time.Time `json:"createTime"`
	DelegatorUserInfo  User      `json:"delegatorUserInfo"`
	DelegateeUserInfos []User    `json:"delegateeUserInfos"`
	DelegateScope      struct {
		Tags []struct {
			TagId   ID     `json:"tagId"`
			TagName string `json:"tagName"`
			TagIcon string `json:"tagIcon"`
		} `json:"tags"`
		Apps []struct {
			AppKey  string `json:"appKey"`
			Title   string `json:"title"`
			IconUrl string `json:"iconUrl"`
		} `json:"apps"`
	} `json:"delegateScope"`
	StartTime      time.Time      `json:"startTime"`
	EndTime        time.Time      `json:"endTime"`
	DelegateStatus DelegateStatus `json:"delegateStatus"`
}

type DelegateType int

const (
	DELEGATE_TYPE_DELEGATION DelegateType = 1
	DELEGATE_TYPE_DELEGATED               = 2
)

func (api DelegateApi) Page(page int, size int, delegateType DelegateType, userId SID) (PageResult[Delegate], error) {
	endpoint := "delegate"
	request := map[string]string{
		"pageNum":  strconv.Itoa(page),
		"pageSize": strconv.Itoa(size),
		"type":     strconv.Itoa(int(delegateType)),
		"userId":   string(userId),
	}
	var result ApiResponse[PageResult[Delegate]]
	err := api.client.get(endpoint, request, &result)
	if err != nil {
		return PageResult[Delegate]{}, err
	}
	return result.Result, nil
}

func (api DelegateApi) Terminate(delegateId ID, userId SID) error {
	endpoint := "delegate/terminate"
	request := map[string]any{
		"delegateId": delegateId,
		"userId":     userId,
	}
	var result ApiResponse[any]
	err := api.client.post(endpoint, request, &result)
	if err != nil {
		return err
	}
	return nil
}

type DelegateCreationRequest struct {
	DelegatesUserIds []SID // 所有委托人的userId
	DelegatedTagIds  []ID
	DelegatedAppKeys []string
	StartTIme        time.Time
	EndTime          time.Time
	userId           SID // 发起委托的用户id
}

func (api DelegateApi) Create(request DelegateCreationRequest) (ID, error) {
	endpoint := "delegate"
	var result ApiResponse[Delegate]
	err := api.client.post(endpoint, request, &result)
	if err != nil {
		return 0, err
	}
	return result.Result.Id, nil
}
