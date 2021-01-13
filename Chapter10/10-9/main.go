package main

import (
	"fmt"
	"time"
)

func GetFood1(c chan string) {
	time.Sleep(3 * time.Second)
	close(c)
}

func GetFood2(c chan string) {
	time.Sleep(3 * time.Second)
	c <- "清蒸鱼好了。"
}

func GetFood3(c chan string) {
	time.Sleep(2 * time.Second)
	c <- "烤生蚝好了。"
}

func main() {
	c1 := make(chan string)
	c2 := make(chan string)
	c3 := make(chan string)

	go GetFood1(c1)
	go GetFood2(c2)
	go GetFood3(c3)

	select {
	case r := <-c1:
		fmt.Println(r)
	case r := <-c2:
		fmt.Println(r)
	case r := <-c3:
		fmt.Println(r)
	default:
		fmt.Println("菜还没有好。")
	}


	select {
	case r:=<-c1:
		fmt.Println(r)
	case r:=<-c2:
		fmt.Println(r)
	case r:=<-c3:
		fmt.Println(r)
		break
	}

}
