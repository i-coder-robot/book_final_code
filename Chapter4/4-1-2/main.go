package main

import "fmt"

func main() {
	foods:=make([]string,6,8)
	foods[0]="红烧肉"
	foods[1]="清蒸鱼"
	foods[2]="溜大虾"
	foods[3]="蒸螃蟹"
	foods[4]="鲍鱼粥"
	fmt.Println(foods[3])

	last := foods[len(foods)-1]
	fmt.Println(last)
}
