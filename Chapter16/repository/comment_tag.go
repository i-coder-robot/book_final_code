package repository

import "github.com/i-coder-robot/book_final_code/Chapter16/model"

type CommentTagRepo struct {
	DB model.DataBase
}

func (t *CommentTagRepo) GetCommentTagList(hotelId string) []model.CommentTag {
	var tagList []model.CommentTag
	t.DB.MyDB.Where("hotel_id=?", hotelId).Find(&tagList)
	return tagList
}
