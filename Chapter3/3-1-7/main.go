package main

import "fmt"

func main() {
	var list = [3]string{"米饭","馒头","包子"}
	list[0]="饺子"
	fmt.Println(list)
}
