package repository

import (
	"github.com/i-coder-robot/book_final_code/Chapter16/model"
)

type OrderSeatRepo struct {
	DB model.DataBase
}

func (s *OrderSeatRepo) OrderSeatOp(o model.OrderSeat) string {
	s.DB.MyDB.Save(o)
	return o.OrderSeatId
}
