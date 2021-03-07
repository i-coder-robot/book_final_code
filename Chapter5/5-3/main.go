package main

import "fmt"

func main() {
	i := 3
	switch i {
	case 0:
		fmt.Printf("红烧肉")
	case 1:
		fmt.Printf("清蒸鱼")
	case 2:
		fmt.Printf("溜大虾")
	case 3:
		fmt.Printf("蒸螃蟹")
	case 4:
		fmt.Printf("鲍鱼粥")
	default:
		fmt.Printf("再等等")
	}

	// 编译不通过
	//prices := [...]int16{32, 68, 96, 153, 198, 77, 100}
	//switch 32 + 68 {
	//case prices[0], prices[1]:
	//	fmt.Println("0 or 1")
	//case prices[2], prices[3]:
	//	fmt.Println("2 or 3")
	//case prices[4], prices[5], prices[6]:
	//	fmt.Println("4 or 5 or 6")
	//}


	prices := [...]int16{32, 68, 96, 153, 198, 77, 100}
	switch prices[1] {
	case 32, 68:
		fmt.Println("0 or 1")
	case 96, 153:
		fmt.Println("2 or 3")
	case 198, 77, 100:
		fmt.Println("4 or 5 or 6")
	}
	// 编译不通过
	//prices := [...]int16{32, 68, 96, 153, 198, 77, 100}
	//switch prices[1] {
	//case 32, 68,96:
	//	fmt.Println("0 or 1 or 2")
	//case 96, 153:
	//	fmt.Println("2 or 3")
	//case 198, 77, 100:
	//	fmt.Println("4 or 5 or 6")
	//}


	prices = [...]int16{32, 68, 96, 153, 198, 77, 100}
	switch prices[1] {
	case prices[0], prices[1],prices[2]:
		fmt.Println("0 or 1 or 2")
	case prices[2], prices[3]:
		fmt.Println("2 or 3")
	case prices[4], prices[5], prices[6]:
		fmt.Println("4 or 5 or 6")
	}

}