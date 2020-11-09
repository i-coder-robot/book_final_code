package model

// 猜你喜欢
type Guess struct {
	//GuessId
	GuessId   string `json:"guess_id"`
	// 图片的地址
	Src       string `json:"src"`
	// 展示标题
	Title     string `json:"title"`
	// 描述
	Desc      string `json:"desc"`
	//优惠价格
	GoodPrice string `json:"goodPrice"`
	//原始价格
	Price     string `json:"price"`
	//售出分数
	SoldNum   int32  `json:"soldNum"`
}