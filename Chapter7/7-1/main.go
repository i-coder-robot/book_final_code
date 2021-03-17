package main

import "fmt"

//方法注释
func Fav() {
	fmt.Println("我喜欢吃生蚝")
}

func add(x, y int) (result int) {
	result = x + y
	return
}

func main() {

	Fav()
}
