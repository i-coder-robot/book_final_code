package repository

import (
	"book-code/Chapter13/13-4/model"
)

type MarketRepo struct {
	DB model.DataBase
}

func (m *MarketRepo) GetMarketInfo(hotelId string) model.Market  {
	var market model.Market
	m.DB.MyDB.Where("hotel_id=?",hotelId).Find(&market)
	return market
}
