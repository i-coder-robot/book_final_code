package dp_service

import (
	"book-code/Chapter13/13-4/model"
	"book-code/Chapter13/13-4/repository"
)

type DiscountService struct {
	Repo *repository.Discount
}

func (d *DiscountService) ListDiscounts() []model.Discount{
	itemsLeft:= d.Repo.ListDiscounts(1)
	itemsRight:= d.Repo.ListDiscounts(2)
	itemsLeft = append(itemsLeft, itemsRight...)
	return itemsLeft
}
