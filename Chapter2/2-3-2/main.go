package main

import (
	"bytes"
	"fmt"
	"strings"
)

func main() {

	name := "peter"
	phone := "13000000000"
	user := name + " " + phone
	fmt.Println(user)

	user1 := strings.Join([]string{"hello", "world"}, " ")
	fmt.Println(user1)

	user2 := fmt.Sprintf("%s:%s", name, phone)
	fmt.Println(user2)

	user3 := bytes.Buffer{}
	user3.WriteString("hello")
	user3.WriteString(" world")
	fmt.Println(user3.String())

}
