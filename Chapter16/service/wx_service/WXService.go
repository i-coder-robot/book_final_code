package wx_service

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/i-coder-robot/book_final_code/Chapter16/res"
	"github.com/spf13/viper"
	"net/http"
)

// 这个函数以 code 作为输入, 返回调用微信接口得到的对象指针和异常情况
func WXLogin(code string) (*res.WXLoginResponse, error) {
	url := "https://api.weixin.qq.com/sns/jscode2session?appid=%s&secret=%s&js_code=%s&grant_type=authorization_code"

	// 这里的appId和secret是在微信公众平台上获取的
	url = fmt.Sprintf(url, viper.GetString("wx_app_id"), viper.GetString("wx_secret"), code)

	// 创建http get请求
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// 解析http请求中body 数据到我们定义的结构体中
	wxResp := res.WXLoginResponse{}
	decoder := json.NewDecoder(resp.Body)
	if err := decoder.Decode(&wxResp); err != nil {
		return nil, err
	}

	// 判断微信接口是否返回一个异常情况
	if wxResp.ErrCode != 0 {
		return nil, errors.New(fmt.Sprintf("ErrCode:%s  ErrMsg:%s", wxResp.ErrCode, wxResp.ErrMsg))
	}

	return &wxResp, nil
}
