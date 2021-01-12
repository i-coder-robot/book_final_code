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

func (c Chef) Cook(name string) string  {
	return c.Name+"：做好了 "+name
}

func (c Chef) FavCook(name string) string {
	return c.Name+"：这是我的拿手菜"+name+"，做好了。"
}

func (c *Chef) GetName() string{
	return c.Name
}

func main() {
	li := Chef{
		Name: "李师傅",
		Age:  25,
		Honor:Honor{},
		Trainee: nil,
	}

	liFav:=li.FavCook
	r:=liFav("葱烧海参")
	fmt.Printf("%s",r)

	wang := Chef{
		Name:    "王师傅",
		Age:     18,
		Honor:   Honor{},
		Trainee: nil,
	}
	fmt.Println(wang.GetName())

	fmt.Println((&li).Cook("红烧肉"))
}
