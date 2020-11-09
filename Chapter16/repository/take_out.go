package repository

import "github.com/i-coder-robot/book_final_code/Chapter16/model"

type TakeOutItemRepo struct {
	DB model.DataBase
}

func (t *TakeOutItemRepo) GetTakeOutListByHotelId(hotelId string) []model.TakeOutItem {
	var takeOutItemList []model.TakeOutItem
	t.DB.MyDB.Table("hotel").Joins("JOIN hotel_food_category on hotel.hotel_id=hotel_food_category.hotel_id").Joins("JOIN take_out on hotel_food_category.hotel_food_category_id=take_out.hotel_food_category_id").Where("hotel.hotel_id=?", hotelId).Select("hotel.hotel_name,hotel_food_category.*,take_out.*").Find(&takeOutItemList)
	return takeOutItemList
}
