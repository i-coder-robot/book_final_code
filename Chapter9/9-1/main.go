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

func (c Chef) Cook() bool {
	fmt.Println(c.Name + "饭菜做好了")
	return true
}
func (c Chef) FavCook(foodName string) bool {
	fmt.Println(c.Name + "的拿手菜" + foodName + "做好了")
	return true
}

func Say(i interface{}) {
	fmt.Printf("(%v, %T)\n", i, i)
}

func main() {
	li := Chef{
		Name: "李师傅",
		Age:  28,
	}
	li.FavCook("红烧肉")
	li.Cook()

	var i interface{}
	Say(i)
	i = 77
	Say(i)
	i = "Go语言学习"
	Say(i)

}
