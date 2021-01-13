package main

import "fmt"

func ShowName(names []string) {
	for _, name := range names {
		fmt.Println("这是：" + name)
	}
}

func ShowName2(names []string) {
	for idx, name := range names {
		if name == "蒜蓉粉丝扇贝" {
			names[idx] = "鲍鱼粥"
		}
		fmt.Println("这是：" + name)
	}
}

func main() {
	list := []string{
		"红烧肉", "清蒸鱼", "溜大虾", "蒸螃蟹", "蒜蓉粉丝扇贝",
	}
	ShowName(list)

	ShowName2(list)
	fmt.Println(list)
}
