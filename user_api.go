package qingflowapi

type UserApi struct {
	client Client
}

func NewUserApi(client Client) UserApi {
	return UserApi{client: client}
}

type User struct {
	AreaCode         string `json:"areaCode"`
	BeingActive      bool   `json:"beingAtive"`
	BeingDisabled    bool   `json:"beingDisabled"`
	CustomDepartment []ID   `json:"customDepartment"`
	CustomRole       []ID   `json:"customRole"`
	Department       []ID   `json:"department"`
	Email            string `json:"email"`
	HeadImg          string `json:"headImg"`
	MobileNum        string `json:"mobileNum"`
	Name             string `json:"name"`
	Role             []ID   `json:"role"`
	UserId           string `json:"userId"`
	OptionId         ID     `json:"optionId"`
}
