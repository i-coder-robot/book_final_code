package main

import "fmt"

type ChefInterface interface {
	Cook() bool
	FavCook(foodName string) bool
	Hidden(b bool) string
	GetHonor() string
}

type Chef struct {
	Name  string
	Age   int
	Honor string
}

func (c *Chef) Cook() bool {
	fmt.Println(c.Name + "饭菜做好了")
	return true
}

func (c *Chef) FavCook(foodName string) bool {
	fmt.Println(c.Name + "的拿手菜" + foodName + "做好了")
	return true
}

func (c Chef) Hidden(b bool) string {
	if b {
		return c.Name + "有隐藏绝活"
	}
	return ""
}

func (c Chef) GetHonor() string {
	return c.Honor
}

func WorkForDinner(c *Chef, foodsName string) {
	var ci ChefInterface = c
	ci.Cook()
	ci.FavCook(foodsName)
}

func main() {
	li := Chef{
		Name: "李师傅",
		Age:  28,
	}
	WorkForDinner(&li, "红烧肉")
	zhang := Chef{
		Name: "张师傅",
		Age:  25,
	}
	WorkForDinner(&zhang, "盐焗鸡")

	var ci ChefInterface
	fmt.Printf("%T\n", ci)
	fmt.Println(ci == nil)
	ci = new(Chef)
	fmt.Printf("%T\n", ci)
}
