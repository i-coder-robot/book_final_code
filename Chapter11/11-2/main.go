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

	for i := 0; i < len(s); i++ {
		go GetSeat(s[i])
	}

	time.Sleep(5 * time.Second)

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

func GetSeatError(name string) {
	m.Lock()
	fmt.Println(name + " 已经抢到位置。")
	time.Sleep(1 * time.Second)
	fmt.Println(name + " 已经离开。")
	//这个地方是模拟忘记写了，会引发错误，所以注释掉
	//m.Unlock()
}

func GetSeatReLock(name string) {
	m.Lock()
	m.Lock()
	fmt.Println(name + " 已经抢到位置。")
	time.Sleep(1 * time.Second)
	fmt.Println(name + " 已经离开。")
	m.Unlock()
	m.Unlock()
}

func GetSeatByPass(name string, locker sync.RWMutex) {
	locker.Lock()
	fmt.Println(name + " 已经抢到位置。")
	time.Sleep(1 * time.Second)
	fmt.Println(name + " 已经离开。")
	locker.Unlock()
}
