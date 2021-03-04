package main

import "sync"

var mu sync.Mutex
var m map[string]int

func lookup(key string) int {
	mu.Lock()
	defer mu.Unlock()
	return m[key]
}

func main() {
	m := make(map[string]int)
	m["abc"] = 123
	m["m"] = 3
	m["g"] = 80
	lookup("g")
}
