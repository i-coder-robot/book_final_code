package main

import (
	"fmt"
	"reflect"
)

type Person struct {
	Age  int    `json:"name" test:"testname"`
	Name string `json:"age" test:"testage"`
}

func main() {
	p := Person{
		Age:  23,
		Name: "小明",
	}
	refType := reflect.TypeOf(p)
	for i := 0; i < refType.NumField(); i++ {
		field := refType.Field(i)
		if jsonItem, ok := field.Tag.Lookup("json"); ok {
			fmt.Println(jsonItem)
		}
		testItem := field.Tag.Get("test")
		fmt.Println(testItem)
	}
}
