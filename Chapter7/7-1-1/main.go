package main

import "fmt"

func Fav(userName string) {
	fmt.Println(userName + " 喜欢吃生蚝")
}

func Fav2(names []string) []string {
	names[0] = names[0] + "班长"
	return names
}

func Fav3(names [2]string) [2]string {
	names[0] = names[0] + "班长"
	return names
}

func main() {
	Fav("小明")
	s := []string{"小明", "小红"}
	fmt.Println(Fav2(s))
	fmt.Println(s)
	arr := [2]string{"小明", "小红"}
	fmt.Println(Fav3(arr))
	fmt.Println(arr)
}
