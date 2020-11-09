package main

import (
	"fmt"
	"reflect"
)

func main() {
	v := "红烧肉"
	t := reflect.TypeOf(v)
	switch t.Kind() {
	case reflect.Int:
		fmt.Println("int")
	case reflect.String:
		fmt.Println("string")
	}
}
