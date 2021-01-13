package main

import "fmt"

func main() {
	var list [5]string
	fmt.Println(list)

	var list2 = [5]string{
		"红烧肉", "清蒸鱼", "溜大虾", "蒸螃蟹", "蒜蓉粉丝扇贝",
	}
	fmt.Println(list2)

	list3 := [5]string{"红烧肉", "清蒸鱼", "溜大虾", "蒸螃蟹", "蒜蓉粉丝扇贝"}
	fmt.Println(list3)

	list4 :=  [...]string{"红烧肉", "清蒸鱼", "溜大虾", "蒸螃蟹", "蒜蓉粉丝扇贝"}
	fmt.Println(list4)

	var list5 =[3]string{"米饭","馒头","包子"}
	//list5=list4
	fmt.Println(list5)
}
