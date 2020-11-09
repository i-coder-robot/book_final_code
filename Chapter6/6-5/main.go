package main

import "fmt"

func main() {
	var m map[string]int = map[string]int{"红烧肉": 88, "清蒸鱼": 98, "溜大虾": 128, "蒸螃蟹": 198, "鲍鱼粥": 68}
	var prices []*int
	//有问题的
	//for k, v := range m {
	//	 fmt.Printf("k:[%p].v:[%p]\n", &k, &v)
	//	 prices = append(prices, &v)
	//}
	//for _, price := range prices {
	//	 fmt.Println(*price)
	//}

	//正确的
	for k, v := range m {
		fmt.Printf("k:[%p].v:[%p]\n", &k, &v)
		vv := m[k]
		prices = append(prices, &vv)
	}
	for _, price := range prices {
		fmt.Println(*price)
	}
}
