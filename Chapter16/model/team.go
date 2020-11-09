package model

//团购
type Team struct {
	//团购Id
	TeamId string `json:"teamId"`
	//餐馆ID
	HotelId string `json:"hotelId"`
	//餐馆名称
	TeamName string `json:"teamName"`
	//原价
	Price int `json:"price"`
	//团购价格
	TeamPrice int `json:"teamPrice"`
	//销售数量
	SoldNum string `json:"soldNum"`
}


