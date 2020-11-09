package res

import "book-code/Chapter13/13-4/model"

type TeamResp struct {
	model.TeamDetail
	FoodList    []TeamChooseFoodResp
	ChooseItems []TeamChooseItemResp
}
