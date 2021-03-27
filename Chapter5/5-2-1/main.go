package main

import "fmt"

func main() {
	name := "麻辣小龙虾"
	if name == "麻辣小龙虾" {
		fmt.Println("晚上吃麻辣小龙虾")
	}

	cause := "正常下班"
	if cause != "加班" {
		fmt.Println("晚上吃大餐去")
	}

}
