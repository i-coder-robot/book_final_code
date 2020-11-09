package dp_service

import (
	"book-code/Chapter13/13-4/model"
	"book-code/Chapter13/13-4/repository"
)

type RestaurantService struct {
	Repo *repository.RestaurantNavRepo
	ItemRepo *repository.RestaurantTabItemRepo
}

func (r *RestaurantService) ListRestaurantNav(level int) []model.RestaurantNav {
	items:= r.Repo.ListRestaurantNav(level)
	return items
}

func (r *RestaurantService) ListGoodRestaurantTabItem() []model.RestaurantTabItem {
	items:= r.ItemRepo.ListGoodRestaurantTabItem()
	return items
}

