package dp_service

import (
	"book-code/Chapter13/13-4/model"
	"book-code/Chapter13/13-4/repository"
)

type GuessService struct {
	Repo *repository.GuessRepo
}

func (g *GuessService) ListGuess() []model.Guess {
	return g.Repo.ListGuesses()
}