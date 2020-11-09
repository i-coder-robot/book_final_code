package dp_service

import (
	"github.com/i-coder-robot/book_final_code/Chapter16/model"
	"github.com/i-coder-robot/book_final_code/Chapter16/repository"
)

type CommentService struct {
	Repo repository.CommentRepo
}

func (c *CommentService) GetCommentList(hotelId string) []model.Comment {
	return c.Repo.GetCommentList(hotelId)
}
