package main

import "fmt"

func main() {
	var m map[string]int = map[string]int{"红烧肉":88,"清蒸鱼":98, "溜大虾":128, "蒸螃蟹":198, "鲍鱼粥":68}
	fmt.Println(m)
	m["麻婆豆腐"] = 48
	m["水煮鱼"] = 99
	fmt.Println(m)

}
