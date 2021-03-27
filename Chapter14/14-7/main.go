package main

import (
	"fmt"
	"reflect"
)

type IFly interface {
	Fly()
}
type Bird struct {
	Name string
}

func (b *Bird) Fly() {}

func main() {
	bird := &Bird{}
	t := reflect.TypeOf((*IFly)(nil)).Elem()
	refType := reflect.TypeOf(bird)
	fmt.Println(refType.Implements(t))
}
