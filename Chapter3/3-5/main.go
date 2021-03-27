package main

import "fmt"

func main() {
	var list = [5]string{
		"红烧肉", "清蒸鱼", "溜大虾", "蒸螃蟹", "蒜蓉粉丝扇贝",
	}
	message := "我点一个菜" + list[1]
	fmt.Println(message)
}
