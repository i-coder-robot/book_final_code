package dp_service

import (
	"book-code/Chapter13/13-4/model"
	"book-code/Chapter13/13-4/repository"
)

type SuggestFoodService struct {
	Repo *repository.SuggestFoodRepo
}

func (sf *SuggestFoodService) ListSuggestList(level int) []model.Suggest {
	return sf.Repo.ListSuggest(level)
}

func (sf *SuggestFoodService) GetFoodByHotelId(hotelId string) []model.SuggestFood {
	return sf.Repo.GetFoodByHotelId(hotelId)
}
