package repository

import "github.com/i-coder-robot/book_final_code/Chapter16/model"

type CommentRepo struct {
	DB model.DataBase
}

func (c *CommentRepo) GetCommentList(hotelId string) []model.Comment {
	var commentList []model.Comment
	c.DB.MyDB.Where("hotel_id=?", hotelId).Find(&commentList)
	return commentList
}
