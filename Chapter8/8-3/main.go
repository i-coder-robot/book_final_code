package main

import (
	"fmt"
	"time"
)

type Chef struct {
	Name    string //名称
	Age     int    //年龄
	Honor          //荣誉
	Trainee *Chef  //徒弟,这里为了演示，可以认为徒弟有多个，用切片表示
}

type Honor struct {
	Title   string    //获奖名称
	GetTime time.Time //获奖时间
}

func main() {
	li := Chef{
		Name:    "李师傅",
		Age:     25,
		Honor:   Honor{},
		Trainee: nil,
	}
	li.Title = "中华金厨奖"
	fmt.Println(li.Name)

	chef := struct {
		Name string
		Age  int
	}{
		Name: "老王",
		Age:  30,
	}
	fmt.Println(chef)
}
