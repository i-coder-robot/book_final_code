package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	//ctx, cancelFunc := context.WithDeadline(context.Background(),time.Now().Add(10*time.Second))
	//ctx.Done()
	//defer cancelFunc()
	//fmt.Println(ctx)
}
