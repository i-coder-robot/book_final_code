package repository

import "book-code/Chapter13/13-4/model"

type  GuessRepo struct {
	DB model.DataBase
}

func (g *GuessRepo) ListGuesses() []model.Guess {
	var guessList []model.Guess
	g.DB.MyDB.Find(&guessList)
	return guessList
}
