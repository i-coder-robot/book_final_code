package main

import (
	"fmt"
	"time"
)

type Chef struct {
	Name    string //名称
	Age     int    //年龄
	*Honor         //荣誉
	Trainee *Chef  //徒弟,这里为了演示，可以认为徒弟有多个，用切片表示
}

type Honor struct {
	Title   string    //获奖名称
	GetTime time.Time //获奖时间
}

func ShowName(names []string) {
	for _, name := range names {
		fmt.Println("这是：" + name)
	}
}

func main() {
	chef := struct {
		Name string
		Age  int
	}{
		Name: "老王",
		Age:  30,
	}
	fmt.Println(chef)

	list := []string{
		"红烧肉", "清蒸鱼", "熘大虾", "蒸螃蟹", "蒜蓉粉丝扇贝",
	}
	ShowName(list)

}
