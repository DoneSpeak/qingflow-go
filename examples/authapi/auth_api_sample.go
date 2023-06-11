package main

import (
	"fmt"
	"os"

	"github.com/DoneSpeak/qingflow-go"
)

func getCredential() qingflow.Credential {
	wsId := os.Getenv("QINGFLOW_WSID")
	wsSecret := os.Getenv("QINGFLOW_WSSECRET")
	return qingflow.Credential{WsId: wsId, WsSecret: wsSecret}
}

func getToken() qingflow.AccessToken {
	api := qingflow.DefaultClient().Auth()
	cred := getCredential()
	fmt.Println("Cred: ", cred)
	token, err := api.GrantToken(cred)
	if err != nil {
		fmt.Println("Err: ", err)
		return nil
	}
	return token
}

func main() {
	token := getToken()
	fmt.Println("Token: ", token)
}
