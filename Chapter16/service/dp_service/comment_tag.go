package dp_service

import (
	"github.com/i-coder-robot/book_final_code/Chapter16/model"
	"github.com/i-coder-robot/book_final_code/Chapter16/repository"
)

type CommentTagService struct {
	Repo repository.CommentTagRepo
}

func (c *CommentTagService) GetCommentTagList(hotelId string) []model.CommentTag {
	return c.Repo.GetCommentTagList(hotelId)
}
