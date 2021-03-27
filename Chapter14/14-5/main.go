package main

import (
	"fmt"
	"reflect"
)

type Person struct {
	Age  int    `json:"NewAge"`
	Name string `json:"NewName"`
}
type newPerson struct {
	NewAge  int
	NewName string
}

func main() {
	t := Person{
		Age:  23,
		Name: "小明",
	}
	refType := reflect.TypeOf(t)
	refValue := reflect.ValueOf(t)
	newPerson := &newPerson{}
	newValue := reflect.ValueOf(newPerson)
	for i := 0; i < refType.NumField(); i++ {
		field := refType.Field(i)
		newTag := field.Tag.Get("json")
		tValue := refValue.Field(i)
		newValue.Elem().FieldByName(newTag).Set(tValue)
	}
	fmt.Println(newPerson)
}
