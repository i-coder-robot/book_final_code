package dp_service

import (
	"book-code/Chapter13/13-4/model"
	"book-code/Chapter13/13-4/repository"
)

type CommentService struct {
	Repo repository.CommentRepo
}

func (c *CommentService) GetCommentList(hotelId string) []model.Comment {
	return c.Repo.GetCommentList(hotelId)
}
