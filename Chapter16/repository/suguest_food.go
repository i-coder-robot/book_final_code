package repository

import "github.com/i-coder-robot/book_final_code/Chapter16/model"

type SuggestFoodRepo struct {
	DB model.DataBase
}

func (s *SuggestFoodRepo) GetFoodByHotelId(hotelId string) []model.SuggestFood {
	var foods []model.SuggestFood
	s.DB.MyDB.Where("hotel_id=?", hotelId).Find(&foods)
	return foods
}

func (s *SuggestFoodRepo) ListSuggest(level int) []model.Suggest {
	var items []model.Suggest
	model.DB.MyDB.Where("level=?", level).Find(&items)
	return items
}
