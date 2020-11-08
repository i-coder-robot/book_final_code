package main

import (
	"fmt"
	"strings"
)

func main() {
	// 提供的主食如下
	foods4:="米饭，馒头，面条"
	result:=strings.Trim(foods4,"米饭")
	fmt.Println(result)

}
