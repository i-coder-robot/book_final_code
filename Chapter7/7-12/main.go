package main

import "fmt"

func start() {
	fmt.Println("程序开始执行...")
}

func testFood() {
	foods :=[]string{"红烧肉", "清蒸鱼", "溜大虾", "蒸螃蟹", "鲍鱼粥"}
	defer func() {
		 if err:=recover();err!=nil{
			 	fmt.Println(err)
			 }
		 fmt.Println("defer finished")
		 }()
	 fmt.Println(foods[5])
}


func end() {
	fmt.Println("程序执行结束...")
}

func main() {
	start()
	testFood()
	end()

}
