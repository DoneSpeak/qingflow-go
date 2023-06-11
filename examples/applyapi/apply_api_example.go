package main

import (
	"os"
	"fmt"
	"net/http"

	"github.com/DoneSpeak/qingflow-go"
)

func filterApply() {
	BaseUrl := "https://api.qingflow.com"
	appKey := os.Getenv("QINGFLOW_APP_KEY")
	token := qingflow.SimpleAccessToken{AccessToken: os.Getenv("QINGFLOW_TOKEN")}

	HttpClient := http.Client{}
	apiClient := qingflow.Client{BaseUrl: BaseUrl, Token: token, HttpClient: HttpClient}
	api := apiClient.Apply(appKey)

	query := qingflow.ApplyQuery{PageSize: 3, PageNumber: 1}
	result, err := api.Query(query)
	if err != nil {
		return
	}

	fmt.Println("Result: ", result)
}
