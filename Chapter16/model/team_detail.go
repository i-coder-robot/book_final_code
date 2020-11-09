package model

type TeamDetail struct {
	//TeamDetailId
	TeamDetailId string `gorm:"column:team_detail_id" json:"teamDetailId"`
	//团购详细名称
	TeamDetailName string `gorm:"column:team_detail_name" json:"teamDetailName"`
	//团购政策
	Policy string `gorm:"column:policy" json:"policy"`
	//团购提示
	Tips string `gorm:"column:tips" json:"tips"`
}

type TeamAggregation struct {
	TeamDetail
	TeamChooseFood
	TeamChooseItem
}