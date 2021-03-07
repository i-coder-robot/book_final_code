package main

import (
	"fmt"
	"sync"
	"time"
)

func GetSeatNoLock(name string) {
	fmt.Println(name + " 已经抢到位置。")
	time.Sleep(1 * time.Second)
	fmt.Println(name + " 已经离开。")
}

var m sync.Mutex

func GetSeat(name string) {
	m.Lock()
	defer m.Unlock()
	fmt.Println(name + " 已经抢到位置。")
	time.Sleep(1 * time.Second)
	fmt.Println(name + " 已经离开。")
}

func main() {
	s := []string{"老张", "老王", "老李"}
	for i := 0; i < len(s); i++ {
		go GetSeatNoLock(s[i])
	}
	time.Sleep(5 * time.Second)
	fmt.Println("-------------------")
	for i := 0; i < len(s); i++ {
		go GetSeat(s[i])
	}
	time.Sleep(5 * time.Second)
	fmt.Println("-------------------")
	for i := 0; i < len(s); i++ {
		go CheckSeatWithRWLock(s[i])
		go GetSeatWithRWLock(s[i])
	}

	time.Sleep(8 * time.Second)
}

var rw sync.RWMutex

func GetSeatWithRWLock(name string) {
	rw.Lock()
	defer rw.Unlock()
	fmt.Println(name + " 已经抢到位置。")
	time.Sleep(1 * time.Second)
	fmt.Println(name + " 已经离开。")
}
func CheckSeatWithRWLock(name string) {
	rw.RLock()
	defer rw.RUnlock()
	fmt.Println(name + " 查看位置。")
	time.Sleep(1 * time.Second)
}
