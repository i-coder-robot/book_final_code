package main

import (
	"fmt"
	"reflect"
)

type T struct{}

func main() {
	name := "GoodDinner"
	t := &T{}
	param1 := reflect.ValueOf(666)
	param2 := reflect.ValueOf("红烧肉")
	params := []reflect.Value{param1, param2}
	reflect.ValueOf(t).MethodByName(name).Call(params)
}

func (t *T) GoodDinner(a int, b string) {
	fmt.Println("吃顿好的，"+b, a)
}
