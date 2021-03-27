package repository

import (
	"github.com/i-coder-robot/book_final_code/Chapter16/model"
)

type TeamPostOrderRepo struct {
	DB model.DataBase
}

func (t *TeamPostOrderRepo) Save(order model.TeamPostOrder) string {
	tx:=t.DB.MyDB.Begin()
	defer tx.Rollback()
	//err := tx.Save(order).Error
	err := tx.Create(order).Error
	if err==nil{
		tx.Commit()
	}
	return order.TeamPostOrderId
}
