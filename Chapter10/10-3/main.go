package main

import (
	"fmt"
	"time"
)

func Work(){
	fmt.Println("我在工作")
}

func main() {
	fmt.Println("主协程开始运行")
	go Work()
	time.Sleep(1*time.Second)
	fmt.Println("主协程运行技术")
}
