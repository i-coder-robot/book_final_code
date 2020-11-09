package main

import "fmt"

func main() {
	var m map[string]int = map[string]int{"红烧肉": 88, "清蒸鱼": 98, "溜大虾": 128, "蒸螃蟹": 198, "鲍鱼粥": 68}
	fmt.Println(m["红烧肉"])

	i,ok := m["麻辣小龙虾"]
	if !ok {
		fmt.Println("没有对应的键元素对")
	} else{
		fmt.Println(i)
	}
}
