package main

import (
	"fmt"
	"os"
	"reflect"

	qingflowapi "github.com/bioelement/qingflow-go"
)

func filterApply() {
	BaseUrl := "https://api.qingflow.com"
	appKey := os.Getenv("QINGFLOW_APP_KEY")
	token := qingflowapi.SimpleAccessToken{AccessToken: os.Getenv("QINGFLOW_TOKEN")}

	apiClient := qingflowapi.NewClient(BaseUrl, token)
	api := apiClient.Apply(appKey)

	query := qingflowapi.ApplyQuery{PageSize: 3, PageNum: 1}
	result, err := api.Page(query)
	if err != nil {
		fmt.Println("Error", err.Error(), reflect.TypeOf(err))
		return
	}

	fmt.Println("Result: ", result)
}

func main() {
	filterApply()
}
