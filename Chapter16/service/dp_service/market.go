package dp_service

import (
	"book-code/Chapter13/13-4/model"
	"book-code/Chapter13/13-4/repository"
)

type MarketService struct {
	Repo repository.MarketRepo
}

func (m *MarketService) GetMarketInfo(hotelId string) model.Market {
	var market model.Market
	m.Repo.DB.MyDB.Where("hotel_id=?",hotelId).Find(&market)
	return market
}
