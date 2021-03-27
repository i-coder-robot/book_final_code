package main

import "fmt"

func main() {
	foods := []string{"红烧肉", "清蒸鱼", "熘大虾", "蒸螃蟹", "蒜蓉粉丝扇贝"}
	//fmt.Println(foods[5])
	fmt.Println(foods[len(foods)-1])
}
