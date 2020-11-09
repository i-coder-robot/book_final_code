package repository

import "book-code/Chapter13/13-4/model"

type SameCityRepo struct {
	DB model.DataBase
}

func (s *SameCityRepo) ListSameCity() ([]model.SameCity) {
	var sameCities []model.SameCity

	model.DB.MyDB.Find(sameCities)
	return sameCities
}
