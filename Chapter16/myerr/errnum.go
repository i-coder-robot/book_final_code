package myerr

import "fmt"

type ErrNum struct {
	Code int
	Message string
}

func (e *ErrNum) Error() string{
	return e.Message
}

type Err struct {
	ErrNum ErrNum
	Err error
}

func New(num ErrNum, err error) *Err {
	return &Err{
		ErrNum: ErrNum{Code: num.Code, Message: num.Message},
		Err:    err,
	}
}

func (e *Err) Add(message string) Err  {
	e.ErrNum.Message += " " + message
	return *e
}

func (err *Err) AddFormat(format string, args ...interface{}) Err {
	err.ErrNum.Message += " " + fmt.Sprintf(format, args...)
	return *err
}

func (err *Err) Error() string {
	return fmt.Sprintf("Err - code: %d, message: %s, error: %s", err.ErrNum.Code, err.ErrNum.Message, err.Err)
}

func IsErrAccountNotFound(err error) bool {
	code, _ := DecodeErr(err)
	return code == ErrAccountNotFound.Code
}



func DecodeErr(err error) (int, string) {
	if err == nil {
		return OK.Code, OK.Message
	}

	switch typed := err.(type) {
	case *Err:
		return typed.ErrNum.Code, typed.ErrNum.Message
	case *ErrNum:
		return typed.Code, typed.Message
	default:
	}

	return InternalServerError.Code, err.Error()
}
