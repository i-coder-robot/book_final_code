package model

//团购套餐选取
type TeamChooseItem struct{
	//TeamChoseItemId
	TeamChoseItemId string `gorm:"column:team_chose_item_id" json:"choseItemId"`
	//套餐名称（如 干贝虾蟹粥）
	TeamChooseItemName string `gorm:"column:team_choose_item_name" json:"chooseItemName"`
	//提示（如两人份）
	TeamChooseItemTip string `gorm:"column:team_choose_item_tip" json:"chooseItemTip"`
	//团购价格
	TeamPrice int `gorm:"column:team_price" json:"teamPrice"`
}
