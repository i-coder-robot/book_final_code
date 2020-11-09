package main

import "fmt"

func Menu(name string,price int) map[string]int{
	m:=make(map[string]int)
	m[name]=price
	return m
}

func main() {
	m:=Menu("红烧肉",88)
	fmt.Println(m)
}
