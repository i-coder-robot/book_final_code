package repository

import (
	"book-code/Chapter13/13-4/model"
)

type OrderSeatRepo struct {
	DB model.DataBase
}
func (s *OrderSeatRepo) OrderSeatOp(o model.OrderSeat) string {
	s.DB.MyDB.Save(o)
	return o.OrderSeatId
}
