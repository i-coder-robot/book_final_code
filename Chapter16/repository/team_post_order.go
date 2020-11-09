package repository

import (
	"book-code/Chapter13/13-4/model"
)


type TeamPostOrderRepo struct {
	DB model.DataBase
}

func (t *TeamPostOrderRepo) Save(order model.TeamPostOrder) string {
	t.DB.MyDB.Save(order)
	return order.TeamPostOrderId
}
