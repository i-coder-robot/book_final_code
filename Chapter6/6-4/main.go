package main

import "fmt"

func main() {
	var m map[string]int = map[string]int{"红烧肉": 88, "清蒸鱼": 98, "溜大虾": 128, "蒸螃蟹": 198, "鲍鱼粥": 68}

	for k,v :=range m{
		fmt.Printf("菜品：%s,价钱:%d\n",k,v)
	}

}
