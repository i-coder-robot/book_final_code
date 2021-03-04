package main

import (
	"fmt"
	"time"
)

type Honor struct {
	Title   string    //获奖名称
	GetTime time.Time //获奖时间
}
type Chef struct {
	Name    string //名称
	Age     int    //年龄
	Honor          //荣誉
	Trainee *Chef  //徒弟，这里为了演示，可以认为徒弟有多个，用切片表示
}

func (c *Chef) Cook(name string) string {
	return c.Name + "：做好了 " + name + "\n"
}
func (c *Chef) FavCook(name string) string {
	return c.Name + "：这是我的拿手菜" + name + "，做好了。\n"
}

func main() {
	li := &Chef{
		Name:    "李师傅",
		Age:     25,
		Honor:   Honor{},
		Trainee: nil,
	}
	result := li.Cook("红烧肉")
	fmt.Printf("%s", result)
	result = li.FavCook("葱烧海参")
	fmt.Printf("%s", result)

	li2 := Chef{
		Name:    "李师傅",
		Age:     25,
		Honor:   Honor{},
		Trainee: nil,
	}
	liPoint := &li2
	result2 := liPoint.Cook("红烧肉")
	fmt.Printf("%s", result2)
	result2 = liPoint.FavCook("葱烧海参")
	fmt.Printf("%s", result2)

	li3 := Chef{
		Name:    "李师傅",
		Age:     25,
		Honor:   Honor{},
		Trainee: nil,
	}

	result3 := (&li3).Cook("红烧肉")
	fmt.Printf("%s", result3)
	result3 = (&li3).FavCook("葱烧海参")
	fmt.Printf("%s", result3)
}
