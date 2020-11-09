package myerr

var (
	// Common errors
	OK                  = &ErrNum{Code: 0, Message: "OK"}
	InternalServerError = &ErrNum{Code: 30001, Message: "内部错误."}
	ErrBind             = &ErrNum{Code: 30002, Message: "请求信息无法转换成结构体."}
	ErrDatabase         = &ErrNum{Code: 30002, Message: "Database error."}
	ErrValidation       = &ErrNum{Code: 30001, Message: "Validation failed."}
	ErrEncrypt          = &ErrNum{Code: 30101, Message: "Error occurred while encrypting the user password."}
	// user errors
	ErrAccountNotFound = &ErrNum{Code: 50102, Message: "用户不存在."}
	ErrPassword        = &ErrNum{Code: 50103, Message: "密码错误."}
	ErrAccountEmpty    = &ErrNum{Code: 50104, Message: "用户名不能为空."}
	ErrPasswordEmpty   = &ErrNum{Code: 50103, Message: "密码不能为空."}
	ErrMissingHeader   =  &ErrNum{Code:    50104,Message: "Header 不存在"}
	ErrToken   =  &ErrNum{Code:    50105,Message: "生成 Token 错误"}
	PassParamCheck = &ErrNum{Code: 60000, Message: "参数校验通过"}
)
