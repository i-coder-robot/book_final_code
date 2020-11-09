package main

import "fmt"

func main() {
	var m map[string]int = map[string]int{"红烧肉": 88, "清蒸鱼": 98, "溜大虾": 128, "蒸螃蟹": 198, "鲍鱼粥": 68}
	price, ok := m["鲍鱼粥"]
	if !ok {
		fmt.Println("没有对应的键值对")
	} else {
		fmt.Println(price)
	}

	//var errorMap = map[interface{}]int{
	//	 "红烧肉": 88,
	//	[]string{"清蒸鱼"}: 98, // 这里会引发panic
	//	"溜大虾": 128,
	//}
	//fmt.Println(errorMap)
}
