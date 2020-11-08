package main

import "fmt"

func main() {
	foods :=[]string{"红烧肉", "清蒸鱼", "溜大虾", "蒸螃蟹", "蒜蓉粉丝扇贝"}
	price := 1288
	if price==1288 && len(foods)>3{
		fmt.Println("赠送优惠卡三张")
	}


	if price==1288 || len(foods)<3{
		fmt.Println("赠送优惠卡一张")
	}


}
