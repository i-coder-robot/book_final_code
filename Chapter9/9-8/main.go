package main

import "fmt"

func main() {
	var x interface{}
	x = "abc"
	switch x.(type) {
	case string:
		fmt.Println("这是string类型")
	case bool:
		fmt.Println("这是bool类型")
	case int:
		fmt.Println("这是int类型")
	}

}
