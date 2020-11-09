package main

import (
	"fmt"
	"sync"
	"time"
)

//func GetSeat(name string) {
//	fmt.Println(name + " 已经抢到位置。")
//	time.Sleep(1 * time.Second)
//	fmt.Println(name + " 已经离开。")
//}
//
//func main() {
//	s := []string{"老张", "老王", "老李"}
//	for i := 0; i < len(s); i++ {
//		go GetSeat(s[i])
//	}
//	time.Sleep(5 * time.Second)
//}


//var m sync.Mutex
//func GetSeat(name string) {
//	m.Lock()
//	defer m.Unlock()
//	fmt.Println(name + " 已经抢到位置。")
//	time.Sleep(1 * time.Second)
//	fmt.Println(name + " 已经离开。")
//}
//
//func main() {
//	s := []string{"老张", "老王", "老李"}
//	for i := 0; i < len(s); i++ {
//		go GetSeat(s[i])
//	}
//	time.Sleep(5 * time.Second)
//}


var m sync.RWMutex
func GetSeat(name string) {
	m.Lock()
	defer m.Unlock()
	fmt.Println(name + " 已经抢到位置。")
	time.Sleep(1 * time.Second)
	fmt.Println(name + " 已经离开。")
}

func CheckSeat(name string){
	m.RLock()
	defer m.RUnlock()
	fmt.Println(name + " 查看位置。")
	time.Sleep(1 * time.Second)
}

func main() {
	s := []string{"老张", "老王", "老李"}
	for i := 0; i < len(s); i++ {
		go CheckSeat(s[i])
		go GetSeat(s[i])
	}
	time.Sleep(8 * time.Second)
}
