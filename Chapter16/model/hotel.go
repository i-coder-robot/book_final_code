package model

type Hotel struct {
	//餐馆ID
	HotelID string `json:"hotelId"`
	//餐馆名称
	HotelName string `json:"hotelName"`
	// 餐馆图片地址
	Pic string `json:"pic"`
	// 餐馆打分
	Star string `json:"star"`
	//评论数
	Num int `json:"num"`
	//人均消费
	AveragePrice float64 `json:"averagePrice"`
	//口味
	Taste float64 `json:"taste"`
	//环境
	Env float64 `json:"env"`
	//服务
	Service float64 `json:"service"`
	//详细地址 （某路某号某广场2层）
	MapAddr string `json:"mapAddr"`
	//详细地址2 （距地铁几号线某站 A 西北口步行多少米）
	MapAddr2 string `json:"mapAddr2"`
	//类型（潮汕菜/湘菜）
	ShortType string `json:"shortType"`
	//营业时间
	BusinessTime string `json:"businessTime"`
	//榜单排行
	Bang string `json:"bang"`
	//团购列表
	TeamList []Team
	//推荐菜列表
	FoodList []SuggestFood
	//评论tag列表
	CommentTagList []CommentTag
	//评论列表
	CommentList []Comment
	//商场
	Market Market
}
