package repository

import (
	"github.com/i-coder-robot/book_final_code/Chapter16/model"
)

type Discount struct {
	DB model.DataBase
}

func (d *Discount) ListDiscounts(pos int) []model.Discount {
	var discounts []model.Discount
	d.DB.MyDB.Where("pos=?", pos).Find(&discounts)
	return discounts
}
