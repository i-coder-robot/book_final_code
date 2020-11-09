package repository

import "github.com/i-coder-robot/book_final_code/Chapter16/model"

type SameCityRepo struct {
	DB model.DataBase
}

func (s *SameCityRepo) ListSameCity() []model.SameCity {
	var sameCities []model.SameCity

	model.DB.MyDB.Find(sameCities)
	return sameCities
}
