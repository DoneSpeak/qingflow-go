package qingflowapi

import (
	"fmt"
)

type RoleApi struct {
	client Client
}

type Role struct {
	RoleId   ID
	RoleName string
	OptionId ID
}

func (api RoleApi) DeleteRoleUser(roleId ID, userIds []ID) (ID, error) {
	endpoint := fmt.Sprintf("role/%d/user", roleId)
	request := map[string]any{
		"roleId":  roleId,
		"userIds": userIds,
	}
	var result ApiResponse[Role]
	err := api.client.deleteRequest(endpoint, request, &result)
	if err != nil {
		return 0, err
	}
	return result.Result.RoleId, nil
}

func (api RoleApi) AddRoleUser(roleId ID, userIds []ID) (ID, error) {
	endpoint := fmt.Sprintf("role/%d/user", roleId)
	request := map[string]any{
		"roleId":  roleId,
		"userIds": userIds,
	}
	var result ApiResponse[Role]
	err := api.client.post(endpoint, request, &result)
	if err != nil {
		return 0, err
	}
	return result.Result.RoleId, nil
}

func (api RoleApi) DeleteRole(roleId ID) (ID, error) {
	endpoint := fmt.Sprintf("role/%d", roleId)
	var result ApiResponse[Role]
	err := api.client.deleteRequest(endpoint, nil, &result)
	if err != nil {
		return 0, err
	}
	return result.Result.RoleId, nil
}

func (api RoleApi) Update(roleId ID, roleName string) (ID, error) {
	endpoint := fmt.Sprintf("role/%d", roleId)
	request := Role{RoleId: roleId, RoleName: roleName}
	var result ApiResponse[Role]
	err := api.client.post(endpoint, request, &result)
	if err != nil {
		return 0, err
	}
	return result.Result.RoleId, nil
}

func (api RoleApi) Create(roleId ID, roleName string) (ID, error) {
	endpoint := fmt.Sprintf("role")
	request := Role{RoleId: roleId, RoleName: roleName}
	var result ApiResponse[Role]
	err := api.client.post(endpoint, request, &result)
	if err != nil {
		return 0, err
	}
	return result.Result.RoleId, nil
}

func (api RoleApi) GetAll() ([]Role, error) {
	endpoint := fmt.Sprintf("role")
	var result ApiResponse[struct {
		Role []Role `json:"role"`
	}]
	err := api.client.get(endpoint, nil, &result)
	if err != nil {
		return []Role{}, err
	}
	return result.Result.Role, nil
}

func (api RoleApi) GetRoleUser(roleId ID) ([]User, error) {
	endpoint := fmt.Sprintf("role")
	var result ApiResponse[struct {
		UserList []User `json:"userList"`
	}]
	err := api.client.get(endpoint, nil, &result)
	if err != nil {
		return []User{}, err
	}
	return result.Result.UserList, nil
}
