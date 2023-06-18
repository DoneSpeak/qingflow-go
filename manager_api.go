package qingflowapi

import (
	"fmt"
)

type ManagerApi struct {
	client Client
}

func NewManagerApi(client Client) ManagerApi {
	return ManagerApi{client: client}
}

/*
删除系统超级管理员
仅「超管accessToken」拥有调用权限。（权限组accessToken无法调用）
*/
func (api ManagerApi) DeleteSuper(ids []ID) error {
	endpoint := "manager/super"
	request := map[string]any{
		"managerUserIds": ids,
	}
	var result ApiResponse[any]
	err := api.client.deleteRequest(endpoint, request, &result)
	if err != nil {
		return err
	}
	return nil
}

/*
获取系统管理员列表
仅「超管accessToken」拥有调用权限。（权限组accessToken无法调用）
*/
func (api ManagerApi) GetSuper() ([]ID, error) {
	endpoint := "manager/super"
	var result ApiResponse[struct {
		ManagerUserIds []ID `json:"managerUserIds"`
	}]
	err := api.client.get(endpoint, nil, &result)
	if err != nil {
		return []ID{}, err
	}
	return result.Result.ManagerUserIds, nil
}

/*
增加超级管理员
仅专有轻流支持；
仅「超管accessToken」拥有调用权限。（权限组accessToken无法调用）

@return amount of the result
*/
func (api ManagerApi) SetSuper(userIds []SID) (int, error) {
	endpoint := "manager/super"
	request := map[string]any{
		"managerUserIds": userIds,
	}
	var result ApiResponse[struct {
		ResultAmount int    `json:"resultAmount"`
		ResultFormat string `json:"resultFormat"`
	}]
	err := api.client.post(endpoint, request, &result)
	if err != nil {
		return 0, err
	}
	return result.Result.ResultAmount, nil
}

func (api ManagerApi) DeleteSub(userIds []SID) ([]SID, error) {
	endpoint := "manager/sub"
	request := map[string]any{
		"subManagerIds": userIds,
	}
	var result ApiResponse[struct {
		SubManagerIds []SID `json:"subManagerIds"`
	}]
	err := api.client.deleteRequest(endpoint, request, &result)
	if err != nil {
		return []SID{}, err
	}
	return result.Result.SubManagerIds, nil
}

type SubManagerUpdationRequest struct {
	SubmanagerName   string     `json:"submanagerName"`    // 子管理员名称
	AuthMemberConfig AuthMember `json:"auth_memberConfig"` // 子管理员成员配置。
	AuthAppConfig    struct {
		BeingAuthApp    bool `json:"beingAuthApp"`    // 应用权限是否开启，true：开启
		BeingEditApp    bool `json:"beingEditApp"`    // 编辑应用是否开启，true：开启。
		BeingAddApp     bool `json:"beingAddApp"`     // 添加应用是否开启，true：开启。
		BeingDelApp     bool `json:"beingDelApp"`     // 删除应用是否开启，true：开启。
		BeingDataManage bool `json:"beingDataManage"` // 数据管理是否开启，true：开启。
	} `json:"auth_app_config"`
	AuthScope struct {
		AppKeys  []string `json:"appKeys"`  // 应用key。
		DeshKeys []string `json:"deshKeys"` // 仪表盘key。
		TagIds   []ID     `json:"tagIds"`   // 应用包ID。
	} `json:"auth_scope"` // 应用/门户权限范围
	AuthContactConfig struct {
		BeingAuthContact bool `json:"beingAuthContact"` // 通讯录权限是否开启，true：开启。
	} `json:"authContactConfig"` // 通讯录权限配置
}

type SubManagerCreationRequest SubManagerUpdationRequest

/*
更新子管理员配置

仅支持专有轻流；
仅「超管accessToken」拥有调用权限。（权限组accessToken无法调用）

注意事项
- 不传/传null，如"userIds":null时，不更新现有配置；
- 传空字符串，如"userIds":" "时，删除全部成员；
- 传具体值，如"userIds":[1,2,3]时，更新成员为：1、2、3。
*/
func (api ManagerApi) UpdateSub(subManagerId ID, request SubManagerUpdationRequest) (ID, error) {
	endpoint := fmt.Sprintf("manager/sub/%s", subManagerId)
	var result ApiResponse[struct {
		SubManagerId ID `json:"subManamerId"`
	}]
	err := api.client.put(endpoint, request, &result)
	if err != nil {
		return 0, err
	}
	return result.Result.SubManagerId, nil
}

type SubManagerConfig struct {
	SubManagerId     SID    `json:"subManagerId"`   // 子管理员ID。
	SubmanagerName   string `json:"submanagerName"` // 子管理员名称
	AuthMemberConfig struct {
		Users []struct {
			UserId   ID     `json:"userId"`
			UserName string `json:"userName"`
		} `json:"users"`
		Departs []struct {
			DeptId   ID     `json:"deptId"`
			DeptName string `json:"deptName"`
		} `json:"departs"`
		Roles []struct {
			RoleId   ID     `json:"roleId"`
			RoleName string `json:"roleName"`
		} `json:"roles"`
		BeingSubDepart bool `json:"beingSubDepart"` // 是否动态包含已选部门下的子部门。（如果部门里有数据，则为必填）
		AppAuthConfig  struct {
			BeingAuthApp     bool `json:"beingAuthApp"`
			BeingEditApp     bool `json:"beingEditApp"`
			BeingAddApp      bool `json:"beingAddApp"`
			BeingDelApp      bool `json:"beingDelApp"`
			BeingDataManager bool `json:"beingDataManager"`
			AuthScope        struct {
				Apps []struct {
					AppKey  string `json:"appKey"`
					AppName string `json:"appName"`
				} `json:"apps"`
				Dashs []struct {
					DashKey  string `json:"dashKey"`
					DashName string `json:"dashName"`
				} `json:"dashs"`
				Tags []struct {
					TagId   ID     `json:"tagId"`
					TagName string `json:"tagName"`
				} `json:"tags"`
			} `json:"auth_scope"`
		} `json:"app_auth_config"`
		AuthContactConfig struct {
			BeingAuthContact bool `json:"beingAuthContact"`
		} `json:"authContactConfig"`
	} `json:"authMemberConfig"` // 子管理员权限配置
}

/*
获取工作区子管理员配置

仅「超管accessToken」拥有调用权限。（权限组accessToken无法调用）
*/
func (api ManagerApi) GetAllSubManager() ([]SubManagerConfig, error) {
	endpoint := "manager/sub"
	var result ApiResponse[struct {
		Submanagers []SubManagerConfig `json:"submanagers"`
	}]
	err := api.client.get(endpoint, nil, &result)
	if err != nil {
		return []SubManagerConfig{}, err
	}
	return result.Result.Submanagers, nil
}

func (api ManagerApi) CreateSubManager(request SubManagerCreationRequest) (ID, error) {
	endpoint := "manager/sub"
	var result ApiResponse[struct {
		SubManagerId ID `json:"subManagerId"`
	}]
	err := api.client.post(endpoint, request, &result)
	if err != nil {
		return 0, err
	}
	return result.Result.SubManagerId, nil
}
