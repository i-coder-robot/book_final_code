package dp_service

import (
	"github.com/hashicorp/go-uuid"
	"github.com/i-coder-robot/book_final_code/Chapter16/handler/param"
	"github.com/i-coder-robot/book_final_code/Chapter16/model"
	"github.com/i-coder-robot/book_final_code/Chapter16/repository"
)

type OrderSeatService struct {
	Repo *repository.OrderSeatRepo
}

func (os *OrderSeatService) OrderSeatOp(p param.OrderSeat) string {
	//order.HotelId-- 到数据库 查看是否有这个团购优惠
	//防止黑客捏造优惠
	//order.time 是否超选择的范围
	//order.Mobile 格式是否正确
	//数据格式要求
	//order.roomtype 检查预订类型是否合法
	//成功预订以后，要发送短信
	id, _ := uuid.GenerateUUID()
	o := model.OrderSeat{
		OrderSeatId: id,
		HotelID:     p.HotelID,
		Mobile:      p.Mobile,
		Name:        p.Name,
		Sex:         p.Sex,
		Message:     p.Message,
	}

	return os.Repo.OrderSeatOp(o)
}
