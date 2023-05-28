package main

import (
	"encoding/json"
	"fmt"
)

type Creature struct {
	Lifetime int `json:"lifetime"`
	Age      int `json:"age"`
}

type Response[R any] struct {
	Code   int `json:"code"`
	Result R   `json:"result"`
}

func main() {
	var r Response[Creature]
	data := `{"code": 1, "result": {"lifetime": 100, "age": 12}}`
	json.Unmarshal([]byte(data), &r)
	fmt.Println("Result: ", r.Result)

	bytes, err := json.Marshal(&r)
	if err != nil {
		fmt.Println("Err: ", err)
		return
	}
	fmt.Println("Result: ", string(bytes))
}
