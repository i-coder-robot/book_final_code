package main

import "fmt"

func Total(prices ...int) int{
	result:=0
	for _,val:=range prices {
		result+=val
	}
	return result
}


func main() {
	fmt.Println(Total(88,98,128,198,68))
}
