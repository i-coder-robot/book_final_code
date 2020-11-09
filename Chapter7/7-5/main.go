package main

import (
	"fmt"
	"net/http"
)

func main() {
	var m map[string]int = map[string]int{"红烧肉": 88, "清蒸鱼": 98, "溜大虾": 128, "蒸螃蟹": 198, "蒜蓉粉丝扇贝": 68}
	v, ok := m["麻辣小龙虾"]
	if !ok {
		fmt.Println("没有对应的键值对")
	} else {
		fmt.Println(v)
	}
	res, err := http.Get("http://www.baidu.com")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(res)

}
