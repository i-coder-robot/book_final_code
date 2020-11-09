package model



//首页导航
type Nav struct{

	NavId string `json:"navId"`
	//图片地址
	Src string `json:"src"`
	//导航标题
	Title string `json:"title"`
}
