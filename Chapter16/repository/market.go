package repository

import (
	"github.com/i-coder-robot/book_final_code/Chapter16/model"
)

type MarketRepo struct {
	DB model.DataBase
}

func (m *MarketRepo) GetMarketInfo(hotelId string) model.Market {
	var market model.Market
	m.DB.MyDB.Where("hotel_id=?", hotelId).Find(&market)
	return market
}
