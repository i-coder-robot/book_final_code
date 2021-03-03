package main

type ICook interface {
	Cook(name string)
	Buy()
	Eat()
}

type Chef struct {
}

func (c Chef) Cook(name string) {

}

func (c Chef) Buy() {

}

func (c Chef) Eat() {

}

func main() {

}
