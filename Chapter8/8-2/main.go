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

type Demo struct {
	A string
	C string
	B string
}

func main() {

	d:=Demo{
		A: "a",
		B: "b",
	}
	fmt.Println(d)

	wang :=Chef{"王师傅",25,Honor{},nil}
	fmt.Printf("%s",wang.Name)
	fmt.Println((&wang).Name)

	li := Chef{
		Name: "李师傅",
		Age:  25,
		Honor:Honor{},
		Trainee: nil,
	}
	fmt.Println(li.Name)


	li2  :=new(Chef)
	*li2 = Chef{
		Name: "李师傅",
		Age:  25,
		Honor:Honor{},Trainee: nil,
	}
	fmt.Println(li2.Name)
}
