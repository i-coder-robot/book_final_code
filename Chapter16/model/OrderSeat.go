package model

//订座
type OrderSeat struct {
	//OrderSeatId
	OrderSeatId string `json:"orderSeatId"`
	//餐馆ID
	HotelID     string `json:"hotelId"`
	//订座人数
	Persons     int    `json:"persons"`
	//预定时间
	DateTime    int64  `json:"datetime"`
	//联系手机号
	Mobile      int    `json:"mobile"`
	//房间类型
	RoomType    int    `json:"room_Type"`
	//联系人名称
	Name        string `json:"name"`
	//联系人性别
	Sex         int    `json:"sex"`
	//留言
	Message     string `json:"message"`
}