package main

import "time"

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

}
