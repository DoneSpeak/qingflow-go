package main

import (
	"fmt"
	"net/http"

	"github.com/donespeak/qingflow-go"
)

func filterApply() {
	BaseUrl := "https://api.qingflow.com"
	appKey := "1f2d6b89"
	token := qingflow.SimpleAccessToken{AccessToken: "45c689b3-c888-4636-b7ce-3c2342712609"}
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
