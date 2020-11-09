package repository

import "book-code/Chapter13/13-4/model"

type MeItemRepo struct {
	DB model.DataBase
}
func (item *MeItemRepo) ListMe() ([]model.MeItem) {
	var items []model.MeItem
	item.DB.MyDB.Find(&items)
	return items
}