package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup
	fmt.Println("可以点菜了吗？")
	wg.Add(1)
	go func() {
		defer wg.Done()
		fmt.Println("老张没没到")
		time.Sleep(1 * time.Second)
		fmt.Println("老张到了")
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		fmt.Println("老王还没到")
		time.Sleep(500 * time.Millisecond)
		fmt.Println("老王到了")
	}()
	fmt.Println("等所有人到齐 ")
	wg.Wait()
	fmt.Println("所有人到齐了，可以点菜了!")

	var once sync.Once
	for i:=0;i<5;i++{
		once.Do(func() {
			fmt.Println("春节快乐。")
		})
	}

}
