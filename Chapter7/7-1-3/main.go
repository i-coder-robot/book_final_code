package main

import "fmt"

func Desc(foodType string, foodName string) {
	fmt.Println("这是" + foodType + "的" + foodName)
}

func main() {
	Desc("卤菜", "葱烧海参")
}
