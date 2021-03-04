package main

import (
	"errors"
	"fmt"
)

type WrapError struct {
	msg string
	err error
}

func (e *WrapError) Error() string {
	return e.msg
}

func (e *WrapError) Unwrap() error {
	return e.err
}

type ErrorString struct {
	s string
}

func (e *ErrorString) Error() string {
	return e.s
}

func main() {
	//err1:=errors.New("My Error 1")
	//err11:=fmt.Errorf("err11:[%w]",err1)
	//err111:=fmt.Errorf("err111: [%w]",err11)
	//e1 := fmt.Errorf("【%w】", err111)
	//fmt.Println(e1)

	err2 := errors.New("My Error 2")
	e2 := fmt.Errorf("err2:【%w】", err2)
	e22 := fmt.Errorf("err22:【%w】", e2)
	fmt.Println(e2)
	fmt.Println(errors.Unwrap(e2))
	fmt.Println(errors.Unwrap(errors.Unwrap(e22)))

	err3 := errors.New("My Error 3")
	e3 := fmt.Errorf("err3:【%w】", err3)
	e33 := fmt.Errorf("err33:【%w】", e3)

	fmt.Println(errors.Is(e33, e3))
	fmt.Println(errors.Is(e33, err3))

	var targetErr *ErrorString
	err4 := fmt.Errorf("new error:[%w]", &ErrorString{s: "target err"})
	fmt.Println(errors.As(err4, &targetErr))
}
