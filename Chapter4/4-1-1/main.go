package main

import "fmt"

func main() {
	var foods =[5]string{"红烧肉", "清蒸鱼", "溜大虾", "蒸螃蟹", "鲍鱼粥"}
	var foodsSlice []string = foods[0:3]
	fmt.Println(foodsSlice)
	fmt.Println(foods[1:4])
	fmt.Println(foods[:2])

	var foodsSlice2 =[]string{"红烧肉", "清蒸鱼", "溜大虾", "蒸螃蟹", "鲍鱼粥"}
	foodsSlice5:= foodsSlice2[1:4]
	fmt.Println(len(foodsSlice5))
	fmt.Println(cap(foodsSlice5))
	foodSlice4 :=[]string{"红烧肉"}
	fmt.Println(foodSlice4)
	fmt.Println(fmt.Sprintf("Len:%d",len(foodSlice4)))
	fmt.Println(fmt.Sprintf("Cap:%d",cap(foodSlice4)))
	foodSlice4 = append(foodSlice4, "清蒸鱼")
	fmt.Println(fmt.Sprintf("Len:%d",len(foodSlice4)))
	fmt.Println(fmt.Sprintf("Cap:%d",cap(foodSlice4)))
	foodSlice4 = append(foodSlice4, "溜大虾")
	fmt.Println(fmt.Sprintf("Len:%d",len(foodSlice4)))
	fmt.Println(fmt.Sprintf("Cap:%d",cap(foodSlice4)))
	foodSlice4 = append(foodSlice4, "蒸螃蟹")
	fmt.Println(fmt.Sprintf("Len:%d",len(foodSlice4)))
	fmt.Println(fmt.Sprintf("Cap:%d",cap(foodSlice4)))
	foodSlice4 = append(foodSlice4, "鲍鱼粥")
	fmt.Println(fmt.Sprintf("Len:%d",len(foodSlice4)))
	fmt.Println(fmt.Sprintf("Cap:%d",cap(foodSlice4)))
}
