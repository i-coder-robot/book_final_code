package main

import (
	"fmt"
	"strings"
)

func main() {
	foods3 := "    米饭，馒头，面条    "
	foods3 = strings.Trim(foods3, " ")
	fmt.Println(foods3)

	foods4 := "米饭，馒头，面条"
	result := strings.Trim(foods4, "米饭")
	fmt.Println(result)

}
