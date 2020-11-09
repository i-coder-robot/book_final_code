package repository

import (
	"github.com/i-coder-robot/book_final_code/Chapter16/model"
)

type TeamPostOrderRepo struct {
	DB model.DataBase
}

func (t *TeamPostOrderRepo) Save(order model.TeamPostOrder) string {
	t.DB.MyDB.Save(order)
	return order.TeamPostOrderId
}
