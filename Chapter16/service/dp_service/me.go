package dp_service

import (
	"book-code/Chapter13/13-4/model"
	"book-code/Chapter13/13-4/repository"
)

type MeService struct {
	Repo *repository.MeItemRepo
}

func (m *MeService) ListMe() []model.MeItem {
	return m.Repo.ListMe()
}
