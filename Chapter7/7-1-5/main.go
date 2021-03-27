package main

import "fmt"

func ShowFoods3(foods []string) {
	foods[0] = "蛋炒饭"
	fmt.Println(foods)
}

func main() {
	todayFoods3 := []string{"卤肉饭", "西红柿鸡蛋饭", "肥牛饭"}
	ShowFoods3(todayFoods3)
	fmt.Println(todayFoods3)

}
