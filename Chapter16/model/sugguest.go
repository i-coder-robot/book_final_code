package model

//首页推荐
type Suggest struct {
	//SuggestId

	SuggestId string `json:"suggestId"`
	//关键字
	KeyWord1 string `json:"keyWord1"`
	KeyWord2 string `json:"keyWord2"`
	KeyWord3 string `json:"keyWord3"`
	KeyWord4 string `json:"keyWord4"`
	//图片地址
	Src string `json:"src"`
	//菜品名称
	FoodName string `json:"foodName"`
	//icon图标
	Icon string `json:"icon"`
	//餐馆名称
	HotelName string `json:"hotelName"`
	//参加人数
	JoinPersons string `json:"joinPersons"`
	//价格
	Price string `json:"price"`
	//优惠价格
	GoodPrice string `json:"goodPrice"`
	//距离我多远
	Distance string `json:"distance"`
	//等级分类
	Level int `json:"level"`
}
