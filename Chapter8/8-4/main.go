package main

import (
	"encoding/json"
	"fmt"
)

type Chef struct {
	Name string
	Age  int
}

func main() {

	c := Chef{
		Name: "老王",
		Age:  28,
	}

	marshal, err := json.Marshal(&c)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(marshal))

	var cc Chef
	s := `{"Name":"小李","Age":24}`
	err = json.Unmarshal([]byte(s), &cc)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(cc.Name)
	fmt.Println(cc.Age)
}
