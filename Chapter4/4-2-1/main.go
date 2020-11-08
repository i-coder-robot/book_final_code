package main

import "fmt"

func main() {
	var foods = []string{"红烧肉", "清蒸鱼", "溜大虾", "蒸螃蟹", "鲍鱼粥"}
	fmt.Println(foods)
	foods[0] = "烤生蚝"
	fmt.Println(foods)
}
