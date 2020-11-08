package main

import "fmt"

func main() {
	//var foods = []string{"红烧肉", "清蒸鱼", "溜大虾", "蒸螃蟹", "鲍鱼粥"}
	//fmt.Println(foods)
	//foods = append(foods, "三文鱼")
	//fmt.Println(foods)
	//
	//foods2:=[]string{}
	//foods2 = append(foods2, "红烧肉")
	//foods2 = append(foods2, "清蒸鱼")
	//foods2 = append(foods2, "溜大虾")
	//foods2 = append(foods2, "蒸螃蟹")
	//foods2 = append(foods2, "鲍鱼粥")
	//fmt.Println(foods2)

	//foods :=[]string{"红烧肉", "清蒸鱼", "溜大虾", "蒸螃蟹", "鲍鱼粥"}
	//foods = append(foods, make([]string, 5)...)
	//fmt.Printf("foods长度：%d\n",len(foods))
	//fmt.Printf("foods容量：%d\n",cap(foods))

	//foods :=[]string{"红烧肉", "清蒸鱼", "溜大虾", "蒸螃蟹", "鲍鱼粥"}
	//foods = append(foods[:2], append([]string{"三文鱼"}, foods[2:]...)...)
	//fmt.Println(foods)

	//foods :=[]string{"红烧肉", "清蒸鱼", "溜大虾", "蒸螃蟹", "鲍鱼粥"}
	//foods = append(foods[:4], append(make([]string, 3), foods[4:]...)...)
	//fmt.Println(foods)

	foods2 :=[]string{"米饭","面条","馒头"}
	var foods = []string{"红烧肉", "清蒸鱼", "溜大虾", "蒸螃蟹", "鲍鱼粥"}
	foods = append(foods[:4], append(foods2, foods[4:]...)...)
	fmt.Println(foods)
}
