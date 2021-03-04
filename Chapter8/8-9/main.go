package main

import (
	"fmt"
	"time"
)

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

func (c Chef3) Cook(name string) string {
	return c.Name + "：做好了 " + name + "\n"
}

func (c Chef3) FavCook(name string) string {
	return c.Name + "：这是我的拿手菜" + name + "，做好了。\n"
}

func main() {
	li := &Chef3{
		Name:    "李师傅",
		Age:     26,
		Honor3:  Honor3{},
		Trainee: nil,
	}
	liFav := li.FavCook
	r := liFav("葱烧海参")
	fmt.Printf("%s", r)

}
