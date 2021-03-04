package main

import "fmt"

func Total(prices ...int) int {
	result := 0
	for _, val := range prices {
		result += val
	}
	return result
}

func Total3(prices ...int) int {
	fmt.Printf("%T\n", Total)
	return 0
}

func Total2(prices []int) int {
	fmt.Printf("%T\n", Total2)
	return 0
}

func main() {
	fmt.Println(Total(88, 98, 128, 198, 68))
	Total2([]int{11})
	Total3(11, 22)
}
