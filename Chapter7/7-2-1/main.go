package main

import "fmt"

func Desc(foodType string,foodName string) string{
	return "这是" +foodType+"的"+foodName
}

func main() {
	result := Desc("粤菜","盐焗鸡")
	fmt.Println(result)
}
