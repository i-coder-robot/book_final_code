package main

import "fmt"

func main() {
	var foods = []string{"红烧肉", "清蒸鱼", "溜大虾", "蒸螃蟹", "鲍鱼粥"}
	foods2 := append(foods[:3], foods[4:]...)
	fmt.Println(foods2)
}
