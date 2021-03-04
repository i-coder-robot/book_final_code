package main

import "fmt"

func MakeFood(c chan string) {
	foods := []string{"红烧肉", "清蒸鱼", "溜大虾", "蒸螃蟹", "鲍鱼粥"}
	for _, item := range foods {
		c <- item
	}
	close(c)
}

func main() {
	ch := make(chan string)
	go MakeFood(ch)
	for i := range ch {
		fmt.Println(i + " 菜好了")
	}
	fmt.Println("您的菜，上齐了。")
}
