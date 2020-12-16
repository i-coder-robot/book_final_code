package model

//我的 页面配置项
type MeItem struct {
	//ItemId
	ItemId string `json:"itemId"`
	//图片地址
	Src string `json:"src"`
	//标题
	Title string `json:"title"`
	Type  string `json:"type"`
}
