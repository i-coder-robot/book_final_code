package main

import "fmt"

func main() {
	var foods =[]string{"红烧肉", "清蒸鱼", "溜大虾", "蒸螃蟹", "鲍鱼粥"}
	var other_foods = foods[:]
	fmt.Println(other_foods)

	var foodsSlice2 =[]string{"红烧肉", "清蒸鱼", "溜大虾", "蒸螃蟹", "鲍鱼粥"}
	foodsSlice6:=[]string{"米饭","面条"}
	copy(foodsSlice6,foodsSlice2)
	fmt.Println(foodsSlice2)
	fmt.Println(foodsSlice6)
}
