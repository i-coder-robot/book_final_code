package main

import "fmt"

func first(c chan<- string) {
	c <- "买菜"
	close(c)
}
func second(c1 <-chan string, c2 chan<- string) {
	r := <-c1
	c2 <- r + " 买肉"
	close(c2)
}

func Cook(c <-chan string) {
	for r := range c {
		fmt.Println(r + "\n已经准备好，吃顿好的！")
	}
}

func main() {
	ch1 := make(chan string)
	ch2 := make(chan string)
	go first(ch1)
	go second(ch1, ch2)
	Cook(ch2)
	fmt.Println("洗碗...")
}
