package model

//美食热门排行榜
type RestaurantNav struct {
	//NavId
	NavId string `json:"navId"`
	//标题
	Title string `json:"title"`
	//描述
	Desc string `json:"desc"`
	//图片地址
	Src string `json:"src"`
	//等级
	Level int `json:"level"`
}

type RestaurantTabItem struct {
	ItemId string `json:"itemId"`
	Src string `json:"src"`
	Title string `json:"title"`
	Star float32 `json:"star"`
	Price string `json:"price"`
	Area string `json:"area"`
	Type string `json:"type"`
	Desc string `json:"desc"`
	Team string `json:"team"`
	Quan string `json:"quan"`
}


