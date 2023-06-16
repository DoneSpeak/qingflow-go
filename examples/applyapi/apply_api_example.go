package main

import (
	"fmt"
	"net/http"
	"os"

	qingflowapi "github.com/DoneSpeak/qingflow-go"
)

func filterApply() {
	BaseUrl := "https://api.qingflow.com"
	appKey := os.Getenv("QINGFLOW_APP_KEY")
	token := qingflowapi.SimpleAccessToken{AccessToken: os.Getenv("QINGFLOW_TOKEN")}

	HttpClient := http.Client{}
	apiClient := qingflowapi.Client{BaseUrl: BaseUrl, Token: token, HttpClient: HttpClient}
	api := apiClient.Apply(appKey)

	query := qingflowapi.ApplyQuery{PageSize: 3, PageNum: 1}
	result, err := api.Query(query)
	if err != nil {
		return
	}

	fmt.Println("Result: ", result)
}
