package model

type HotelFoodCategory struct {
	//HotelFoodCategoryId
	HotelFoodCategoryId string `json:"hotelFoodCategoryId"`
	// 分类名称
	HotelFoodCategoryName string `json:"hotelFoodCategoryName"`
	// 餐馆ID
	HotelId string `json:"hotelId"`
}

