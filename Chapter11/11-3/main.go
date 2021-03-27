package main

import (
	"fmt"
	"sync"
)

func main() {
	var buySth string
	var box sync.RWMutex
	sendCond := sync.NewCond(&box)
	recvCond := sync.NewCond(box.RLocker())
	flag := make(chan bool, 5)
	go func() {
		defer func() {
			fmt.Println("已投递，完成~")
			flag <- true
		}()
		box.Lock()
		for buySth == "已投递" {
			sendCond.Wait()
		}
		buySth = "已投递"
		box.Unlock()
		recvCond.Signal()
	}()

	go func() {
		defer func() {
			fmt.Println("货物已经取走")
			flag <- true
		}()
		box.RLock()
		for buySth == "" {
			recvCond.Wait()
		}
		buySth = ""
		box.RUnlock()
		sendCond.Signal()
	}()

	<-flag
	<-flag
}
