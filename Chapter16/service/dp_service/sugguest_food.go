package dp_service

import (
	"github.com/i-coder-robot/book_final_code/Chapter16/model"
	"github.com/i-coder-robot/book_final_code/Chapter16/repository"
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

type SuggestService struct {
	Repo *repository.SuggestRepo
}

func (sf *SuggestService) GetSuggestByLevel(level int) model.Suggest {
	return sf.Repo.ListSuggest(level)
}
