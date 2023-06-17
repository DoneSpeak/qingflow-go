package qingflowapi

import (
	"fmt"
	"strconv"
	"time"
)

type AppApi struct {
	client Client
}

func NewAppApi(client Client) AppApi {
	return AppApi{client: client}
}

type AppAuthType int

const (
	APP_AUTH_WORKSPACE      AppAuthType = 1 // Default
	APP_AUTH_SPECIFIED_USER             = 2
)

type AppUpdationRequest struct {
	// 新增应用所属的应用包，此处只能选择自定义应用包。不可选系统默认应用包
	TagIds      []AppTagId  `json:"tagIds"`
	AppAuth     AppAuthType `json:"appAuth"`
	AuthMembers []struct {
		UserIds []int64 `json:"userIds"`
		DeptIds []int64 `json:"deptIds"`
		RoleIds []int64 `json:"roleIds"`
	} `json:"authMembers"`
	IncludeSubDeparts bool `json:"includeSubDeparts"`
}

type AppCreationRequest struct {
	AppTagUpdateRequest
	// 应用名称，默认为“未命名应用”
	AppName string `json:"appName"`
	// 应用的图标，默认为系统图标
	AppIcon string `json:"appIcon"`
	// 应用的创建人userid，默认为工作区创建者
	UserId string `json:"userId"`
}

func (api AppApi) Create(request AppCreationRequest) (string, error) {
	endpoint := "app"
	var result ApiResponse[struct {
		AppKey string `json:"appKey"`
	}]

	err := api.client.post(endpoint, request, &result)
	if err != nil {
		return "", err
	}
	return result.Result.AppKey, nil
}

func (api AppApi) Delete(appKey string) (string, error) {
	endpoint := fmt.Sprintf("app/%s", appKey)
	var result ApiResponse[struct {
		AppKey string `json:"appKey"`
	}]
	err := api.client.deleteRequest(endpoint, nil, &result)
	if err != nil {
		return "", err
	}
	return result.Result.AppKey, nil
}

func (api AppApi) Update(appKey string, request AppTagUpdateRequest) (string, error) {
	endpoint := fmt.Sprint("app/%s", appKey)
	var result ApiResponse[struct {
		AppKey string `json:"appKey"`
	}]
	err := api.client.post(endpoint, request, &result)
	if err != nil {
		return "", err
	}
	return result.Result.AppKey, nil
}

type App struct {
	AppKey      string
	AppName     string
	AppAuth     string
	AuthMembers []struct {
		Users []struct {
			UserId   string
			UserName string
		}
		Depts []struct {
			DeptId   int
			DeptName string
		}
		Roles []struct {
			RoleId   int
			RoleName string
		}
	}
	Creator struct {
		UserId   string
		NickName string
		HeadImg  string
	}
	CreateTime time.Time
	AppIcon    string
	Tags       []struct {
		TagId   string
		TagName string
	}
	AppPublishStatus int
}

func (api AppApi) Page(appKey string, page int, size int) (PageResult[App], error) {
	endpoint := "apps"
	var result ApiResponse[PageResult[App]]
	params := map[string]string{
		"appKey":   appKey,
		"pageNum":  strconv.Itoa(page),
		"pageSize": strconv.Itoa(size),
	}
	err := api.client.get(endpoint, params, &result)
	if err != nil {
		return PageResult[App]{}, err
	}
	return result.Result, nil
}

type QuestionBaseInfo struct {
	QueId    ID
	QueType  QueType
	QueTitle string
	Required bool
	Options  []struct {
		OptId      ID
		LinkQueIds []ID
		OptValue   string
	}
	SubQuestionBaseInfos []QuestionBaseInfo
	DisplayedQueId       ID
	QlinkerAlias         string
	QlinkerConfig        struct {
		QlinkerAlias string
		QueId        ID
		QueTitle     string
	}
	ReletionType int
}

type AppForm struct {
	QuestionBaseInfos []QuestionBaseInfo
}

func (api AppApi) GetForm(appKey string) (AppForm, error) {
	endpoint := fmt.Sprint("app/%s", appKey)
	var result ApiResponse[AppForm]
	err := api.client.get(endpoint, nil, &result)
	if err != nil {
		return AppForm{}, err
	}
	return result.Result, nil
}

type AppPrintTemplate struct {
	FormTitle string
	PrintTpls []struct {
		PrintKey       string
		PrintName      string
		CustomFileName string
	}
}

// 获取应用word打印模版信息
func (api AppApi) GetPrintTemplate(appKey string) (AppPrintTemplate, error) {
	endpoint := fmt.Sprint("app/%s/print", appKey)
	var result ApiResponse[AppPrintTemplate]
	err := api.client.get(endpoint, nil, &result)
	if err != nil {
		return AppPrintTemplate{}, err
	}
	return result.Result, nil
}

// @param userId 成员外部id，可查询指定成员的可见应用。userId为空时，返回工作区全部应用。
// @param userId不传时无效，获取当前用户可见中收藏的应用
func (api AppApi) GetAll(userId string, favourite bool) ([]App, error) {
	endpoint := "app"
	var result ApiResponse[struct {
		AppList []App
	}]
	params := map[string]string{
		"userId": userId,
	}
	if favourite {
		params["favourite"] = "1"
	}
	err := api.client.get(endpoint, params, &result)
	if err != nil {
		return []App{}, nil
	}
	return result.Result.AppList, nil
}
