package main

import (
	"fmt"
	"sync"
)

var m map[string]int
var  mu sync.Mutex

func lookup(key string) int{
	mu.Lock()
	defer mu.Unlock()
	return m[key]
}


func main() {
	m=make(map[string]int)
	m["abc"]=123
	m["m"]=3
	m["g"]=80
	n:=lookup("g")
	fmt.Println(n)
}
