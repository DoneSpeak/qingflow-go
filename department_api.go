package qingflowapi

import (
	"fmt"
	"strconv"
)

type DepartmentApi struct {
	client Client
}

func NewDepartmentApi(client Client) DepartmentApi {
	return DepartmentApi{client: client}
}

type Department struct {
	DeptId     ID     `json:"deptId"`
	DeptLeader []ID   `json:"deptLeader"`
	Name       string `json:"name"`
	OptionId   ID     `json:"optionId"`
	ParentId   ID     `json:"parentId"`
}

func (api DepartmentApi) DeleteUser(deptId ID, userIds []ID) error {
	endpoint := fmt.Sprintf("department/%s", deptId)
	request := map[string]any{
		"deptId":  deptId,
		"userIds": userIds,
	}
	var result ApiResponse[Department]
	err := api.client.deleteRequest(endpoint, request, &result)
	if err != nil {
		return err
	}
	return nil
}

func (api DepartmentApi) AddUser(deptId ID, userIds []ID) error {
	endpoint := fmt.Sprintf("department/%s", deptId)
	request := map[string]any{
		"deptId":  deptId,
		"userIds": userIds,
	}
	var result ApiResponse[Department]
	err := api.client.post(endpoint, request, &result)
	if err != nil {
		return err
	}
	return nil
}

func (api DepartmentApi) GetAllUser(deptId ID, fetchChild bool) ([]ID, []User, error) {
	endpoint := fmt.Sprintf("department/%s/user", deptId)
	params := map[string]string{}
	if fetchChild {
		params["fetchChild"] = "true"
	}
	var result ApiResponse[struct {
		LeaderIds []ID   `json:"leaderIds"`
		UserList  []User `json:"userList"`
	}]
	err := api.client.get(endpoint, params, &result)
	if err != nil {
		return []ID{}, []User{}, err
	}
	return result.Result.LeaderIds, result.Result.UserList, nil
}

func (api DepartmentApi) Update(dept Department) (ID, error) {
	endpoint := fmt.Sprintf("/department/%s", dept.DeptId)
	var result ApiResponse[Department]
	err := api.client.post(endpoint, dept, &result)
	if err != nil {
		return 0, err
	}
	return result.Result.DeptId, nil
}

func (api DepartmentApi) Create(dept Department) (ID, error) {
	endpoint := fmt.Sprintf("/department", dept.DeptId)
	var result ApiResponse[Department]
	err := api.client.post(endpoint, dept, &result)
	if err != nil {
		return 0, err
	}
	return result.Result.DeptId, nil
}

func (api DepartmentApi) GetAll(deptId ID) ([]Department, error) {
	endpoint := fmt.Sprintf("/department")
	var result ApiResponse[struct {
		Department []Department `json:"department"`
	}]
	params := map[string]string{}

	if deptId > 0 {
		params["deptId"] = strconv.Itoa(int(deptId))
	}

	err := api.client.get(endpoint, params, &result)
	if err != nil {
		return []Department{}, err
	}
	return result.Result.Department, nil
}

func (api DepartmentApi) Delete(deptId ID) (ID, error) {
	endpoint := fmt.Sprintf("/department/%s", deptId)
	var result ApiResponse[Department]
	err := api.client.delete(endpoint)
	if err != nil {
		return 0, err
	}
	return result.Result.DeptId, nil
}

func (api DepartmentApi) GetUndeparted() ([]ID, []User, error) {
	endpoint := "user/undeparted"
	var result ApiResponse[struct {
		LeaderIds []ID   `json:"leaderIds"`
		UserList  []User `json:"userList"`
	}]
	err := api.client.get(endpoint, nil, &result)
	if err != nil {
		return []ID{}, []User{}, err
	}
	return result.Result.LeaderIds, result.Result.UserList, nil
}
