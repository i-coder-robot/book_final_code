package main

import "fmt"

func main() {
	foods := make([]string, 6, 8)
	foods[0] = "红烧肉"
	foods[1] = "清蒸鱼"
	foods[2] = "溜大虾"
	foods[3] = "蒸螃蟹"
	foods[4] = "鲍鱼粥"
	for idx, item := range foods {
		fmt.Println(idx)
		fmt.Println(item)
		fmt.Println("-------")
	}
}
