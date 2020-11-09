package res

type TakeOutItemResp struct {
	//TakeOutId
	TakeOutId string `json:"takeOutId"`
	//餐馆分类ID
	HotelFoodCategoryId string `json:"hotelFoodCategoryId"`
	//菜品名称
	FoodName string `json:"foodName"`
	//菜品图片
	Pic string `json:"pic"`
	//月销售数量
	MonthSoldNum string `json:"monthSoldNum"`
	//好评率
	Zan int `json:"zan"`
	//价格
	Price int `json:"price"`
	//是否为推荐
	IsSuggest int `json:"isSuggest"`
	//折扣价格
	DiscountPrice int `json:"discountPrice"`
	//优惠
	Discount float64 `json:"discount"`
}
