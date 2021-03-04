package main

import (
	"fmt"
)

type Hi func(num string) string

func SayHello(num string, hi Hi) {
	result := hi(num)
	fmt.Println(result)
}

func main() {
	target := "东北食客"

	if target == "东北食客" {
		SayHello("3", func(num string) string {
			return num + "位兄弟，欢迎光临"
		})
	} else {
		SayHello("3", func(num string) string {
			return num + "位客人，欢迎光临"
		})
	}
}
