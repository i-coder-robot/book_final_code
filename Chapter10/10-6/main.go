package main

import "fmt"

func Eat(foodName string, c chan bool) {
	fmt.Println("我正在吃" + foodName)
	c <- true
}

func main() {
	fmt.Println("主协程开始运行")
	c := make(chan bool)
	go Eat("生蚝", c)
	fmt.Println("主协程运行技术")
	<-c
}
