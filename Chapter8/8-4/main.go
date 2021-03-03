package main

import (
	"encoding/json"
	"fmt"
	"time"
)

type Chef struct {
	Name string
	Age  int
	Honor
	Trainee *Chef
}

type Honor struct {
	Title   string    //获奖名称
	GetTime time.Time //获奖时间
}

func Cook(c Chef, name string) string {
	return c.Name + "：做好了 " + name
}

func (c Chef) Cook(name string) string {
	return c.Name + "：做好了 " + name
}

func (c Chef) FavCook(name string) string {
	return c.Name + "：这是我的拿手菜" + name + "，做好了。"
}

func main() {

	c := Chef{
		Name: "老王",
		Age:  28,
	}

	resultFunc := Cook(c, "锅包肉")
	fmt.Println(resultFunc)

	marshal, err := json.Marshal(&c)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(marshal))

	var cc Chef
	s := `{"Name":"小李","Age":24}`
	err = json.Unmarshal([]byte(s), &cc)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(cc.Name)
	fmt.Println(cc.Age)

	li := Chef{
		Name:    "李师傅",
		Age:     25,
		Honor:   Honor{},
		Trainee: nil,
	}
	result := li.Cook("红烧肉")
	fmt.Println(result)
	result = li.FavCook("葱烧海参")
	fmt.Println(result)
}
