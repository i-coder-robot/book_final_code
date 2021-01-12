package main

import "fmt"

func Fav(userName string){
	fmt.Println(userName+" 喜欢吃生蚝")
}
func add(x, y int) (result int) {
	result = x + y
	return
}

func main() {
	Fav("小明")
}
