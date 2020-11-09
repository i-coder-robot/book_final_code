package dp_service

import (
	"github.com/i-coder-robot/book_final_code/Chapter16/model"
	"github.com/i-coder-robot/book_final_code/Chapter16/repository"
)

type MeService struct {
	Repo *repository.MeItemRepo
}

func (m *MeService) ListMe() []model.MeItem {
	return m.Repo.ListMe()
}
