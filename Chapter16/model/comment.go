package model

type Comment struct {
	//评论ID
	CommentId string `json:"commentId"`
	//评论
	Comment string `json:"comment"`
	//评论人
	AccountName string `json:"accountName"`
	//评论人头像
	AccountPic string `json:"accountPic"`
	//评分
	Star int `json:"star"`
	//人均消费
	AveragePerson int `json:"averagePerson"`
	//是否是优质点评
	IsGood int `json:"isGood"`
}