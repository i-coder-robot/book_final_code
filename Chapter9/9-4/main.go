package main

import (
	"fmt"
	"time"
)

type Chef struct {
	Name string
	Age int
	Honor
	Trainee *Chef
}


type Honor struct {
	Title   string    //获奖名称
	GetTime time.Time //获奖时间
}

type ChefInterface interface {
	Cook() bool
	FavCook(foodName string) bool
}
type FarmInterface interface {
	FarmVegeTables(string) string
}
type AdvanceChef interface {
	ChefInterface
	FarmInterface
}

func (c Chef) Cook() string  {
	return c.Name+"饭菜做好了 "}

func ( c Chef)  FavCook(name string) string {
	return c.Name+"：这是我的拿手菜"+name+"，做好了。"
}
func (c *Chef) FarmVegeTables(name string) string{
	return name+"可以收割了。\n"
}

func main() {
	zhang := Chef{
		Name: "张师傅",
		Age:  28,
	}
	fmt.Printf("%s",zhang.FarmVegeTables("豆角"))
	r:=zhang.FavCook("干煸豆角")
	fmt.Println(r)
	r=zhang.Cook()
	fmt.Println(r)
}
