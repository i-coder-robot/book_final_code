package res

import "github.com/i-coder-robot/book_final_code/Chapter16/model"

type TeamResp struct {
	model.TeamDetail
	FoodList    []TeamChooseFoodResp
	ChooseItems []TeamChooseItemResp
}
