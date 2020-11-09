package dp_service

import (
	"github.com/i-coder-robot/book_final_code/Chapter16/model"
	"github.com/i-coder-robot/book_final_code/Chapter16/repository"
)

type GuessService struct {
	Repo *repository.GuessRepo
}

func (g *GuessService) ListGuess() []model.Guess {
	return g.Repo.ListGuesses()
}
