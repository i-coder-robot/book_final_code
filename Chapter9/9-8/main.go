package main

import "fmt"

type Chef struct {
	Name  string
	Age   int
	Honor string
}

type ChefInterface interface {
	GetHonor() string
	Hidden(b bool) string
}

func (c Chef) Hidden(b bool) string {
	if b {
		return c.Name + "有隐藏绝活"
	}
	return ""
}


func (c Chef) GetHonor() string{
	return c.Honor
}


func main() {
	var x interface{}
	x="abc"
	switch x.(type) {
		case string:
			fmt.Println("")
		case bool:
			fmt.Println("")
		case int:
			fmt.Println("")
	}

}
