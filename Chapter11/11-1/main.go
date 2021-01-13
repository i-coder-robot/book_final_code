package main

import "fmt"

var balance int
func Sum(amount int) {
	balance = balance + amount
}
func GetTotal() int {
	return balance
}


func main() {
	for i:=0;i<10;i++{
		go Sum(i)
	}
	fmt.Println(GetTotal())
}
