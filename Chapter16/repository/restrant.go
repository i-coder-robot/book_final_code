package repository

import "github.com/i-coder-robot/book_final_code/Chapter16/model"

type RestaurantNavRepo struct {
	DB model.DataBase
}

func (r *RestaurantNavRepo) ListRestaurantNav(level int) []model.RestaurantNav {
	var items []model.RestaurantNav
	model.DB.MyDB.Where("level=?", level).Find(&items)
	return items
}

type RestaurantTabItemRepo struct {
	DB model.DataBase
}

func (r *RestaurantTabItemRepo) ListGoodRestaurantTabItem() []model.RestaurantTabItem {
	var items []model.RestaurantTabItem
	r.DB.MyDB.Find(&items)
	return items
}

func (r *RestaurantTabItemRepo) ListRestaurantNav(level int) []model.RestaurantNav {
	var items []model.RestaurantNav
	r.DB.MyDB.Where("level=?", level).Find(&items)
	return items
}
