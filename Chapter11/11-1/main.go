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
var m1 sync.RWMutex
var m sync.RWMutex

func GetSeat(name string) {
	m.Lock()
	defer m.Unlock()
	fmt.Println(name + " 已经抢到位置。")
	time.Sleep(1 * time.Second)
	fmt.Println(name + " 已经离开。")
}

func CheckSeat(name string) {
	m.RLock()
	defer m.RUnlock()
	fmt.Println(name + " 查看位置。")
	time.Sleep(1 * time.Second)
}

func GetSeatError(name string){
	m.Lock()
	fmt.Println(name + " 已经抢到位置。")
	time.Sleep(1 * time.Second)
	fmt.Println(name + " 已经离开。")
	//m.Unlock()
}

func GetSeatReLock(name string){
	m.Lock()
	m.Lock()
	fmt.Println(name + " 已经抢到位置。")
	time.Sleep(1 * time.Second)
	fmt.Println(name + " 已经离开。")
	m.Unlock()
	m.Unlock()
}

func GetSeatByPass(name string,locker sync.RWMutex){
	locker.Lock()
	fmt.Println(name + " 已经抢到位置。")
	time.Sleep(1 * time.Second)
	fmt.Println(name + " 已经离开。")
	locker.Unlock()
}

func main() {
	s := []string{"老张", "老王", "老李"}
	//for i := 0; i < len(s); i++ {
	//	go CheckSeat(s[i])
	//	go GetSeat(s[i])
	//}
	//time.Sleep(8 * time.Second)

	//for i := 0; i < len(s); i++ {
	//	go CheckSeat(s[i])
	//	go GetSeatError(s[i])
	//}
	//time.Sleep(8 * time.Second)

	//for i := 0; i < len(s); i++ {
	//	go CheckSeat(s[i])
	//	go GetSeatReLock(s[i])
	//}
	//time.Sleep(8 * time.Second)

	for i := 0; i < len(s); i++ {
		go CheckSeat(s[i])
		go GetSeatByPass(s[i],m)
	}
	time.Sleep(8 * time.Second)
}
