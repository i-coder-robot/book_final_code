package main

import "fmt"

func Greet(name string) string {
	fmt.Println("Hello," + name)
	return "你也好"
}

func main() {

	result := Greet("小明")
	fmt.Println(result)

	var m map[string]int = map[string]int{"红烧肉": 88, "清蒸鱼": 98, "熘大虾": 128, "蒸螃蟹": 198, "蒜蓉粉丝扇贝": 68}
	v, ok := m["麻辣小龙虾"]
	if !ok {
		fmt.Println("没有对应的键值对")
	} else {
		fmt.Println(v)
	}
}
