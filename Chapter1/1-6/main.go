package main

import "fmt"

var msg = "去外面吃点好的。"

func main() {
	msg := "main函数内"
	{
		msg := "在家吃点好的"
		fmt.Println(msg)
	}
	fmt.Println(msg)
}
