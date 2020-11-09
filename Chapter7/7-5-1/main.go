package main

import "fmt"

func Greet(name string) string {
	fmt.Println("Hello," + name)
	return "你也好"

}

func main() {
	result:=Greet("小明")
	fmt.Println(result)
}
