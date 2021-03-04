package main

import "fmt"

type ChefInterface interface {
	GetHonor() string
	Hidden(b bool) string
}
type Chef struct {
	Name  string
	Age   int
	Honor string
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
func (c *Chef) SetHonor(title string) {
	c.Honor = title
}

func main() {
	zhang := Chef{
		Name:  "张师傅",
		Age:   36,
		Honor: "米其林1星",
	}
	fmt.Println(zhang.GetHonor())

	var ci ChefInterface = zhang
	zhang.SetHonor("米其林3星")
	fmt.Println(zhang.GetHonor())
	fmt.Println(ci.GetHonor())

	var zhang2 Chef
	var ci2 ChefInterface
	if ci2 == nil {
		fmt.Println("ci2 is nil")
	}
	ci2 = zhang2

	if ci2 == nil {
		fmt.Println("ci2 is nil,too")
	}
	zhang2.SetHonor("米其林3星")
	fmt.Println(zhang2.GetHonor())
	fmt.Println(ci2.GetHonor())

}
