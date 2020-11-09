package main

import "fmt"

type ChefInterface interface {
	Cook() bool
	FavCook(foodName string) bool
}

type Chef struct {
	Name string
	Age  int
}
func (c *Chef) Cook() bool {
	fmt.Println(c.Name + "饭菜做好了")
	fmt.Printf("%p\n",c)
	return true
}

func (c *Chef) FavCook(foodName string) bool {
	fmt.Println(c.Name + "的拿手菜" + foodName + "做好了")
	fmt.Printf("%p\n",c)
	return true
}
func WorkForDinner(c *Chef,foodsName string){
	var ci ChefInterface=c
	ci.Cook()
	ci.FavCook(foodsName)
}

func main() {
	li := Chef{
		Name: "李师傅",
		Age:  28,
	}
	fmt.Printf("李师傅对象的地址%p\n",&li)
	WorkForDinner(&li,"红烧肉")
	zhang:=Chef{
		Name: "张师傅",
		Age:  25,
	}
	fmt.Printf("张师傅对象的地址%p\n",&zhang)
	WorkForDinner(&zhang,"盐焗鸡")
}
