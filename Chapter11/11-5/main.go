package main

import (
	"context"
	"fmt"
	"math/rand"
	"strconv"
	"sync/atomic"
	"time"
)

func main() {
	Free()

	d := time.Now().Add(100 * time.Millisecond)
	ctx, cancel := context.WithDeadline(context.Background(), d)
	defer cancel()
	select {
	case <-time.After(1 * time.Second):
		fmt.Println("overslept")
	case <-ctx.Done():
		fmt.Println(ctx.Err())
	}
}

func Free() {
	fmt.Println("如果今天有过生日的人，会说出一个数字，这个数字如果和我们的一致，则这顿饭免单")
	secret := rand.Intn(100)
	cxt, cancelFunc := context.WithCancel(context.Background())
	defer cancelFunc()
	for i := 1; i <= 3; i++ {
		var num int32
		r := rand.Intn(100)
		num = int32(r)
		fmt.Printf("我猜的数字是: %d\n", num)
		go choose(func() {
			if atomic.LoadInt32(&num) == int32(secret) {
				fmt.Println("猜对了，给您免单")

			}
		})
	}
	go func() {
		<-cxt.Done()
	}()
	fmt.Println("答案是" + strconv.Itoa(secret))
	fmt.Println("感谢光临，聚餐结束。")
}

func choose(deferFunc func()) {
	defer func() {
		deferFunc()
	}()
	time.Sleep(time.Second * 2)
}
