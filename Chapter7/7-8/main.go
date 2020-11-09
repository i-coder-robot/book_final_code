package main

import (
	"fmt"
	"strconv"
)

type Discount func() float64
type CheckSum func(name string, price float64) float64

func PayOrder(discount Discount) CheckSum {
	var total float64
	return func(name string, price float64) float64 {
		fmt.Println("菜品名称:" + name + "单价:" + strconv.FormatFloat(price, 'f', -1, 64))
		total = total + price
		if discount == nil {
			return total
		}
		return total * discount()
	}
}

func main() {
	f:=PayOrder(func() float64 {
		return 0.8
	})
	result:=f("红烧肉", 88)
	fmt.Println(result)
	result=f("清蒸鱼",98)
	fmt.Println(result)
	result=f("溜大虾", 128)
	fmt.Println(result)
	result=f("蒸螃蟹", 198)
	fmt.Println(result)
	result=f("蒜蓉粉丝扇贝",68)
	fmt.Println(result)

}
