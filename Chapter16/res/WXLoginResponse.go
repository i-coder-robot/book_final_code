package res

type WXLoginResponse struct {
	OpenId string			`json:"openid"`
	SessionKey string		`json:"session_key"`
	UnionId string			`json:"unionid"`
	ErrCode int				`json:"errcode"`
	ErrMsg string 			`json:"errmsg"`
}

