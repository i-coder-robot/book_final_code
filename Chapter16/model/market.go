package model

type Market struct {
	//MarketId
	MarketId string `json:"marketId"`
	//商场名称
	MarketName string `json:"marketName"`
	//商场地址
	Addr string `json:"addr"`
	//类型
	ShortType string `json:"shortType"`
	//评价
	Star int `json:"star"`
	//商场图片地址
	Pic string `json:"pic"`
}
