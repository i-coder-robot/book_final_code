package dp_service

import (
	"errors"
	"github.com/i-coder-robot/book_final_code/Chapter16/model"
	"github.com/i-coder-robot/book_final_code/Chapter16/repository"
)

type HotelService struct {
	Repo            *repository.HotelRepo
	TeamRepo        *repository.TeamRepo
	SuggestFoodRepo *repository.SuggestFoodRepo
	CommentTagRepo  *repository.CommentTagRepo
	CommentRepo     *repository.CommentRepo
	MarketRepo      *repository.MarketRepo
}

func (h *HotelService) GetHotelDetailByID(id string) (*model.Hotel, error) {
	if id == "" {
		return nil, errors.New("参数错误！")
	}
	hotel := h.Repo.GetHotelById(id)
	if &hotel == nil {
		return nil, errors.New("餐馆查询错误！")
	}
	teamList := h.TeamRepo.GetTeamListByHotelId(id)
	hotel.TeamList = teamList
	foodList := h.SuggestFoodRepo.GetFoodByHotelId(id)
	hotel.FoodList = foodList
	tagList := h.CommentTagRepo.GetCommentTagList(id)
	hotel.CommentTagList = tagList
	commentList := h.CommentRepo.GetCommentList(id)
	hotel.CommentList = commentList
	market := h.MarketRepo.GetMarketInfo(id)
	hotel.Market = market
	return &hotel, nil
}
