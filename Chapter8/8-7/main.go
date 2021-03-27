package main

import (
	"fmt"
	"time"
)

type Chef2 struct {
	Name string
	Age  int
	Honor2
	Trainee *Chef2
}

type Honor2 struct {
	Title   string    //获奖名称
	GetTime time.Time //获奖时间
}

func (c *Chef2) Cook(name string) string {
	return c.Name + "：做好了 " + name + "\n"
}

func (c *Chef2) FavCook(name string) string {
	return c.Name + "：这是我的拿手菜" + name + "，做好了。\n"
}

func (c *Chef3) GetName() string {
	return c.Name
}

type Honor3 struct {
	Title   string    //获奖名称
	GetTime time.Time //获奖时间
}

type Chef3 struct {
	Name string
	Age  int
	Honor3
	Trainee *Chef3
}

func (c *Chef3) Cook(name string) string {
	return c.Name + "：做好了 " + name + "\n"
}

func (c *Chef3) FavCook(name string) string {
	return c.Name + "：这是我的拿手菜" + name + "，做好了。\n"
}

func main() {
	wang := Chef2{
		Name:    "王师傅",
		Age:     23,
		Honor2:  Honor2{},
		Trainee: nil,
	}

	result := wang.Cook("番茄炒蛋")
	fmt.Printf("%s", result)
	result = wang.FavCook("糖心鲍鱼")
	fmt.Printf("%s", result)

	zhao := Chef3{
		Name:    "赵师傅",
		Age:     26,
		Honor3:  Honor3{},
		Trainee: nil,
	}
	fmt.Printf("%s", zhao.Cook("蛋炒饭"))
	fmt.Printf("%s", zhao.FavCook("小炒肉"))
}
