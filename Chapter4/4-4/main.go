package main

import "fmt"

func main() {
	foods := []string{"红烧肉", "清蒸鱼", "溜大虾", "蒸螃蟹", "鲍鱼粥"}
	foods2 := make([]*string, len(foods))
	for i, value := range foods {
		foods2[i] = &value
	}
	fmt.Println(foods[0], foods[1], foods[2], foods[3], foods[4])
	fmt.Println(*foods2[0], *foods2[1], *foods2[2], *foods2[3], *foods2[4])

	//解决上面代码的错误，可以参照书中的例子
	//foods := []string{"红烧肉", "清蒸鱼", "溜大虾", "蒸螃蟹", "鲍鱼粥"}
	//foods2 := make([]*string, len(foods))
	//for i, _ := range foods {
	//	foods2[i] = &foods[i]
	//}
	//fmt.Println(foods[0], foods[1], foods[2], foods[3], foods[4])
	//fmt.Println(*foods2[0], *foods2[1], *foods2[2], *foods2[3], *foods2[4])
}
