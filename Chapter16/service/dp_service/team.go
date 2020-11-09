package dp_service

import (
	"github.com/i-coder-robot/book_final_code/Chapter16/model"
	"github.com/i-coder-robot/book_final_code/Chapter16/repository"
	"github.com/i-coder-robot/book_final_code/Chapter16/res"
)

type TeamService struct {
	Repo *repository.TeamRepo
}

func (t *TeamService) GetTeamListByHotelId(hotelId string) []model.Team {
	return t.Repo.GetTeamListByHotelId(hotelId)
}

func (t *TeamService) GetTeamDetail(teamId string) res.TeamResp {
	detail := t.Repo.GetTeamDetail(teamId)
	item := convert2TeamResp(detail)
	return item
}

func convert2TeamResp(items []model.TeamAggregation) res.TeamResp {
	var resp res.TeamResp
	resp.TeamDetail = items[0].TeamDetail
	for _, item := range items {
		f := &item.TeamChooseFood
		resp.FoodList = append(resp.FoodList, Convert2TeamChooseFoodResp(f))
		i := &item.TeamChooseItem
		resp.ChooseItems = append(resp.ChooseItems, TeamChooseItem(i))
	}
	return resp
}
func TeamChooseItem(item *model.TeamChooseItem) res.TeamChooseItemResp {
	r := res.TeamChooseItemResp{
		TeamChoseItemId:    item.TeamChoseItemId,
		TeamChooseItemName: item.TeamChooseItemName,
		TeamChooseItemTip:  item.TeamChooseItemTip,
		TeamPrice:          item.TeamPrice,
	}
	return r
}
func Convert2TeamChooseFoodResp(food *model.TeamChooseFood) res.TeamChooseFoodResp {
	r := res.TeamChooseFoodResp{}
	r.TeamChooseFoodId = food.TeamChooseFoodId
	r.TeamChooseFoodName = food.TeamChooseFoodName
	return r
}
