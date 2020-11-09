package dp_service

import (
	"github.com/i-coder-robot/book_final_code/Chapter16/model"
	"github.com/i-coder-robot/book_final_code/Chapter16/repository"
)

type RestaurantService struct {
	Repo     *repository.RestaurantNavRepo
	ItemRepo *repository.RestaurantTabItemRepo
}

func (r *RestaurantService) ListRestaurantNav(level int) []model.RestaurantNav {
	items := r.Repo.ListRestaurantNav(level)
	return items
}

func (r *RestaurantService) ListGoodRestaurantTabItem() []model.RestaurantTabItem {
	items := r.ItemRepo.ListGoodRestaurantTabItem()
	return items
}
