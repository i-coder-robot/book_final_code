package repository

import "github.com/i-coder-robot/book_final_code/Chapter16/model"

type MeItemRepo struct {
	DB model.DataBase
}

func (item *MeItemRepo) ListMe() []model.MeItem {
	var items []model.MeItem
	item.DB.MyDB.Find(&items)
	return items
}
