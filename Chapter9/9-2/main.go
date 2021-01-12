package main

import "fmt"

//定义一个鸭子接口
type Duck interface {
	Cook()
}

//吊炉烤鸭
type DiaoLu struct{}

func (dl *DiaoLu) Cook() {
	fmt.Println("今天去吃,吊炉烤鸭")
}

//焖炉烤鸭
type MenLu struct{}

func (ml *MenLu) Cook() {
	fmt.Println("今天去吃,焖炉烤鸭")
}

func CookDuck(d Duck) {
	d.Cook()
}

func main() {
	fmt.Println("duck typing")
	var d DiaoLu
	var m MenLu

	CookDuck(&d)
	CookDuck(&m)
}
