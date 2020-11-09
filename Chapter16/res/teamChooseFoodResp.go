package res

//团购套餐选取（如烧腊2选1）
type TeamChooseFoodResp struct {
	//TeamChooseFoodId
	TeamChooseFoodId string `gorm:"column:team_choose_food_id" json:"teamChooseFoodId"`
	//名称
	TeamChooseFoodName string `gorm:"column:team_choose_food_name" json:"teamChooseName"`
}
