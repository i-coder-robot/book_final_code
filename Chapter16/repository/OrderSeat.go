package repository

import (
	"github.com/i-coder-robot/book_final_code/Chapter16/model"
	"log"
)

type OrderSeatRepo struct {
	DB model.DataBase
}

func (s *OrderSeatRepo) OrderSeatOp(o model.OrderSeat) string {
	err := s.DB.MyDB.Save(o).Error
	if err != nil {
		log.Println(err)
	}
	return o.OrderSeatId
}
