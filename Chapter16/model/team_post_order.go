package model

//团购下单
type TeamPostOrder struct {
	TeamPostOrderId string `json:"teamPostOrderId"`
	//团购详情ID
	TeamDetailId string `json:"teamDetailId"`
	//支付价格
	RealPrice int `json:"realPrice"`
	//数量
	Quantity int `json:"quantity"`
	//下单人手机号
	Mobile string `json:"mobile"`
	//下单人名称
	Name string `json:"name"`
	//下单人性别
	Sex int `json:"sex"`
	//附加消息
	Message string `json:"message"`
}
