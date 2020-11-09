package utils

import (
	"book-code/Chapter13/13-4/model"
	"book-code/Chapter13/13-4/myerr"
	"github.com/go-playground/validator/v10"
	"golang.org/x/crypto/bcrypt"
)

// 给文本加密
func Encrypt(source string) (string, error) {
	hashedBytes, err := bcrypt.GenerateFromPassword([]byte(source), bcrypt.DefaultCost)
	return string(hashedBytes), err
}

// 比较密码
func Compare(hashedPassword, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}

// 验证字段是否有效.
func Validate(m model.Account) error {
	validate := validator.New()
	return validate.Struct(m)
}

func CheckParam(accountName,password string) myerr.Err {
	if accountName == "" {
		return myerr.New(*myerr.ErrValidation, nil).Add("用户名为空.")
	}

	if password == "" {
		return myerr.New(*myerr.ErrValidation, nil).Add("密码为空.")
	}
	return myerr.Err{ErrNum: *myerr.PassParamCheck, Err: nil}
}

