package res

type TakeOutRsp struct {
	HotelName string `json:"hotelName"`
	CategoryList []HotelFoodCategoryResp
	ItemList []TakeOutItemResp
}