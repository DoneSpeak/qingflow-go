package main

import (
	"fmt"
	"os"

	qingflowapi "github.com/bioelement/qingflow-go"
)

func getCredential() qingflowapi.Credential {
	wsId := os.Getenv("QINGFLOW_WSID")
	wsSecret := os.Getenv("QINGFLOW_WSSECRET")
	return qingflowapi.Credential{WsId: wsId, WsSecret: wsSecret}
}

func getToken() qingflowapi.AccessToken {
	api := qingflowapi.DefaultClient(nil).Auth()
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
