package repository

import (
	"book-code/Chapter13/13-4/model"
)

type TeamRepo struct {
	DB model.DataBase
}

func (t *TeamRepo) GetTeamListByHotelId(hotelId string) []model.Team{
	var teamList []model.Team
	t.DB.MyDB.Where("hotel_id=?",hotelId).Find(&teamList)
	return teamList
}

func (t *TeamRepo) GetTeamDetail(teamDetailId string) []model.TeamAggregation {
	var teamAggregations  []model.TeamAggregation
	t.DB.MyDB.Table("team_detail").Joins("JOIN team_choose_food on team_detail.team_detail_id=team_choose_food.team_detail_id").Joins("JOIN team_choose_item on team_choose_food.team_choose_food_id=team_choose_item.team_choose_food_id").Select("team_detail.*,team_choose_food.*,team_choose_item.*").Scan(&teamAggregations).Debug()

	return teamAggregations
}

