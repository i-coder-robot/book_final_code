package repository

import "book-code/Chapter13/13-4/model"

type HotelRepo struct {
	DB model.DataBase
}

func (h *HotelRepo) GetHotelById(hotelId string) model.Hotel{
	var hotel model.Hotel
	h.DB.MyDB.Where("hotel_id = ?",hotelId).First(&hotel)
	return hotel
}
