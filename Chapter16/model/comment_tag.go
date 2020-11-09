package model

type CommentTag struct {
	//TagID
	TagId   string `json:"tagId"`
	//Tag名称
	TagName string `json:"tagName"`
	//Tag数
	TagNum  int    `json:"tagNum"`
	//餐馆ID
	HotelId string `json:"hotelId"`
}

