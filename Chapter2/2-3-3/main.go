package main

import "fmt"

func main() {
	message := "hello Go world"
	fmt.Println(message)
	fmt.Println("\t"+message)

	foods := "清蒸鱼\n红烧肉\n鱼香肉丝"
	fmt.Println(foods)

	foods2 := "您选择的菜品:\n\t清蒸鱼\n\t红烧肉\n\t鱼香肉丝"
	fmt.Println(foods2)


}
