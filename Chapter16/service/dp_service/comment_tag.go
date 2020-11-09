package dp_service

import (
	"book-code/Chapter13/13-4/model"
	"book-code/Chapter13/13-4/repository"
)

type CommentTagService struct {
	Repo repository.CommentTagRepo
}

func (c *CommentTagService) GetCommentTagList(hotelId string) []model.CommentTag {
	return c.Repo.GetCommentTagList(hotelId)
}
