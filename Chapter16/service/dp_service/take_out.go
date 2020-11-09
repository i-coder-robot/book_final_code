package dp_service

import (
	"github.com/i-coder-robot/book_final_code/Chapter16/model"
	"github.com/i-coder-robot/book_final_code/Chapter16/repository"
	"github.com/i-coder-robot/book_final_code/Chapter16/res"
)

type TakeOutService struct {
	Repo *repository.TakeOutItemRepo
}

func Convert2HotelFoodCategoryResp(c model.HotelFoodCategory) res.HotelFoodCategoryResp {

	r := res.HotelFoodCategoryResp{
		HotelFoodCategoryId:   c.HotelFoodCategoryId,
		HotelFoodCategoryName: c.HotelFoodCategoryName,
		HotelId:               c.HotelId,
	}
	return r
}

func convert2TakeOutItemResp(t *model.TakeOut) res.TakeOutItemResp {
	r := res.TakeOutItemResp{
		TakeOutId:           t.TakeOutId,
		HotelFoodCategoryId: t.HotelFoodCategoryId,
		FoodName:            t.FoodName,
		Pic:                 t.Pic,
		MonthSoldNum:        t.MonthSoldNum,
		Zan:                 t.Zan,
		Price:               t.Price,
		IsSuggest:           t.IsSuggest,
		DiscountPrice:       t.DiscountPrice,
		Discount:            t.Discount,
	}
	return r
}

func (to *TakeOutService) GetTakeOutListByHotelId(hotelId string) res.TakeOutRsp {
	items := to.Repo.GetTakeOutListByHotelId(hotelId)
	var takeoutRsp res.TakeOutRsp
	if len(items) > 0 {
		takeoutRsp.HotelName = items[0].HotelName
		for _, item := range items {
			c := model.HotelFoodCategory{
				HotelFoodCategoryId:   item.HotelFoodCategory.HotelFoodCategoryId,
				HotelFoodCategoryName: item.HotelFoodCategoryName,
				HotelId:               item.HotelID,
			}
			takeoutRsp.CategoryList = append(takeoutRsp.CategoryList, Convert2HotelFoodCategoryResp(c))
			t := model.TakeOut{
				TakeOutId:           item.TakeOutId,
				HotelFoodCategoryId: item.HotelFoodCategory.HotelFoodCategoryId,
				FoodName:            item.FoodName,
				Pic:                 item.TakeOut.Pic,
				MonthSoldNum:        item.MonthSoldNum,
				Zan:                 item.Zan,
				Price:               item.Price,
				IsSuggest:           item.IsSuggest,
				DiscountPrice:       item.DiscountPrice,
				Discount:            item.Discount,
			}
			takeoutRsp.ItemList = append(takeoutRsp.ItemList, convert2TakeOutItemResp(&t))
		}
	}
	return takeoutRsp
}
