package main

import "fmt"

func main() {
	var m map[string]int = map[string]int{"红烧肉":88,"清蒸鱼":98, "溜大虾":128, "蒸螃蟹":198, "鲍鱼粥":68}
	fmt.Println(m)

	foodsMap := make(map[string]int)
	foodsMap["红烧肉"] = 88
	foodsMap["清蒸鱼"] = 98
	foodsMap["溜大虾"] = 128
	foodsMap["蒸螃蟹"] = 198
	foodsMap["鲍鱼粥"] = 68
	fmt.Println(foodsMap)
}
