// @Title
// @Description
// @Author
package qingflowapi

import (
	"strconv"
	"time"
)

type AppTagApi struct {
	client Client
}

func NewAppTagApi(client Client) AppTagApi {
	return AppTagApi{client: client}
}

type AppTagId int64

type AppType int

const (
	EMPTY_APP    AppType = 0
	SOLUTION_APP         = 1
)

// 可见权限, , 2:
type TagAuthType string

const (
	TAG_AUTH_WORKSPACE      TagAuthType = "1" // 工作区可查看（默认值）
	TAG_AUTH_SPECIFIED_USER             = "2" // 指定用户可见
)

type AppTagUpdateRequest struct {
	TagId   int         `json:"tagId"`
	TagName string      `json:"tagName"` // 应用包名称（1-100字符，可含空格）
	TagAuth TagAuthType `json:"tagAuth"`
	// 应用包图标，格式为链接或指定格式字符串（如下）
	// 默认值"{"icon":null,"iconColor":"qing-orange","iconName":null}"（应用包首字+橙色底）
	TagIcon string `json:"tagIcon"`
	// tagAuth为2时填写，默认为（超/系统）管理员可见；tagAuth为1时填写无效
	AuthMembers []struct {
		UserIds           []string `json:"userIds"`           // 用户id列表
		DeptIds           []string `json:"deptIds"`           // 部门id列表
		RoleIds           []string `json:"roleIds"`           // 角色id列表
		IncludeSubDeparts bool     `json:"includeSubDeparts"` // 是否动态包含已选部门下的子部门，默认true
	} `json:"authMembers"`
}

type AppTagCreationRequest struct {
	AppTagUpdateRequest

	// 应用包的创建人userid（必须为工作区系统管理员）
	// 不填默认为工作区创建者
	UserId string  `json:"userId"`
	Type   AppType `json:"type"`
	// 解决方案ID, type为1时必填（填0时此参数无效）,只支持免费的、单应用包的解决方案
	SolutionKey string `json:"solutionKey"`
	IncludeDate bool   `json:"includeDate"` // 是否包含示例数据（type为1时有效，非必填，默认值为true）
}

type AppTag struct {
	TagId       AppTagId `json:"tagId"`
	TagName     string   `json:"tagName"`
	TagAuth     string   `json:"tagAuth"`
	AuthMembers []struct {
		Users []struct {
			UserId   string `json:"userId"`
			UserName string `json:"userName"`
		} `json:"users"`
		Depts []struct {
			DeptId   int    `json:"deptId"`
			DeptName string `json:"deptName"`
		} `json:"depts"`
		Roles []struct {
			RoleId   int    `json:"roleId"`
			RoleName string `json:"roleName"`
		} `json:"roles"`
		Creator struct {
			UserId   string `json:"userId"`
			NickName string `json:"nickName"`
			HeadImg  string `json:"headImg"`
		} `json:"creator"`
		CreateTime time.Time `json:"createTime"`
		TagIcon    string    `json:"tagIcon"`
		AppList    []struct {
			AppKey  string `json:"appKey"`
			AppName string `json:"appName"`
		} `json:"app_list"`
		DashList []struct {
			DashKey  string `json:"dashKey"`
			DashName string `json:"dashName"`
		} `json:"dashList"`
	} `json:"authMembers"`
}

/**
 * 仅「超管accessToken」拥有调用权限。（权限组accessToken无法调用）
 *
 * @return app tag id
 */
// TODO test
func (api AppTagApi) Create(request AppTagCreationRequest) (AppTagId, error) {
	endpoint := "tag"
	var result ApiResponse[AppTag]
	err := api.client.post(endpoint, request, &result)
	if err != nil {
		return 0, err
	}
	return result.Result.TagId, nil
}

// TODO Test
func (api AppTagApi) Update(request AppTagUpdateRequest) (AppTagId, error) {
	endpoint := "tag"
	var result ApiResponse[AppTag]
	err := api.client.post(endpoint, request, &result)
	if err != nil {
		return 0, err
	}
	return result.Result.TagId, nil
}

type AppTagDeletionRequest struct {
	ClearAll bool       `json:"clearAll"`
	TagIds   []AppTagId `json:"tagIds"`
}

/*
本接口支持批量删除应用包，但仅返回删除成功的应用包id
（若接口报错表示所有应用包均未删除成功）

若clearAll参数传false，应用包中的应用将不被删除,
为「无所属应用包」状态，您仍然可以在工作区内搜索到对应应用。
*/
func (api AppTagApi) Delete(request AppTagDeletionRequest) ([]AppTagId, error) {
	endpoint := "tag"
	var result ApiResponse[struct {
		TagIds []AppTagId
	}]
	err := api.client.deleteRequest(endpoint, request, &result)
	if err != nil {
		return []AppTagId{}, err
	}
	return result.Result.TagIds, nil
}

func (api AppTagApi) GetPage(tagId string, page int, size int) (PageResult[AppTag], error) {
	endpoint := "tag"
	var result ApiResponse[PageResult[AppTag]]
	params := map[string]string{
		"pageNum":  strconv.Itoa(page),
		"pageSize": strconv.Itoa(size),
	}
	err := api.client.get(endpoint, params, &result)
	if err != nil {
		return PageResult[AppTag]{}, err
	}
	return result.Result, nil
}

type AppTagAllResult struct {
	AppList []struct {
		AppKey  string `json:"appKey"`
		AppName string `json:"appName"`
	} `json:"app_list"`
	TagList []struct {
		TagName string `json:"tagName"`
		TagId   string `json:"tagId"`
		TagIcon string `json:"tagIcon"`
	} `json:"tag_list"`
	DashList []struct {
		DashKey  string `json:"dashKey"`
		DashName string `json:"dashName"`
	} `json:"dashList"`
}

// @param favourite 可查询指定成员的可见且收藏的应用列表，userId不传时无效
func (api AppTagApi) GetAll(userId string, favourite bool) (AppTagAllResult, error) {
	endpoint := "tags"
	params := map[string]string{
		"userId": userId,
	}
	if favourite {
		params["favourite"] = "1"
	}
	var result ApiResponse[AppTagAllResult]
	err := api.client.get(endpoint, params, &result)
	if err != nil {
		return AppTagAllResult{}, err
	}
	return result.Result, nil
}
