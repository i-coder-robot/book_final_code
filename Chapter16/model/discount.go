package model

type Discount struct {
	//自增ID
	Id string `json:"id"`
	//DiscountId
	DiscountId string `json:"discountId"`
	//优惠icon地址
	DiscountItemIcon string `json:"discountItemIcon"`
	//优惠项目src
	DiscountItemSrc string `json:"discountItemSrc"`
	//优惠标题
	DiscountItemTitle string `json:"discountItemTitle"`
	//优惠餐馆
	DiscountItemHotel string `json:"discountItemHotel"`
	//优惠价格
	DiscountItemGoodPrice int32 `json:"discountItemGoodsPrice"`
	//原始价格
	DiscountItemPrice int32 `json:"discountItemPrice"`
	//折扣
	Discount string `json:"discount"`
	//TODO 最后做
	Pos int `json:"pos"`
}
