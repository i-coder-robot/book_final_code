package dp_service

import (
	"github.com/i-coder-robot/book_final_code/Chapter16/model"
	"github.com/i-coder-robot/book_final_code/Chapter16/repository"
)

type MarketService struct {
	Repo repository.MarketRepo
}

func (m *MarketService) GetMarketInfo(hotelId string) model.Market {
	var market model.Market
	m.Repo.DB.MyDB.Where("hotel_id=?", hotelId).Find(&market)
	return market
}
