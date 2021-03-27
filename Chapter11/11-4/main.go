package main

import (
	"fmt"
	"sync/atomic"
)

func main() {
	var num1 uint64
	atomic.AddUint64(&num1, 70)
	fmt.Println(num1)
	atomic.AddUint64(&num1, ^uint64(23))
	fmt.Println(num1)

	var num2 uint64
	num2 = 5
	ok := atomic.CompareAndSwapUint64(&num2, 5, 50)
	fmt.Println(ok)
	fmt.Println(num2)
	ok = atomic.CompareAndSwapUint64(&num2, 40, 50)
	fmt.Println(ok)
	fmt.Println(num2)

	var num3 uint64
	num3 = 10
	num := atomic.LoadUint64(&num3)
	fmt.Println(num)

	var num4 uint64
	num4 = 1
	atomic.StoreUint64(&num4, 28)
	fmt.Println(num4)
}
