package main

import "fmt"

func ShowFoods(foods [3]string) {
	//试图修改第一个元素
	foods[0] = "蛋炒饭"
	fmt.Println(foods)
}

func ShowFoods2(foods *[3]string) {
	foods[0] = "蛋炒饭"
	fmt.Println(foods)
}

func ShowFoods3(foods []string) {
	foods[0] = "蛋炒饭"
	fmt.Println(foods)
}

func main() {
	//定义并初始化一个数组
	todayFoods := [3]string{"卤肉饭", "西红柿鸡蛋饭", "肥牛饭"}
	//传递给函数，并试图在函数体内修改这个数组内容
	ShowFoods(todayFoods)
	fmt.Println(todayFoods)

	todayFoods2 := &[3]string{"卤肉饭", "西红柿鸡蛋饭", "肥牛饭"}
	ShowFoods2(todayFoods2)
	fmt.Println(todayFoods2)

	todayFoods3 := []string{"卤肉饭", "西红柿鸡蛋饭", "肥牛饭"}
	ShowFoods3(todayFoods3)
	fmt.Println(todayFoods3)
}