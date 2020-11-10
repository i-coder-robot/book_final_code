package dp_service

import (
	"errors"
	"github.com/hashicorp/go-uuid"
	"github.com/i-coder-robot/book_final_code/Chapter16/handler/param"
	"github.com/i-coder-robot/book_final_code/Chapter16/model"
	"github.com/i-coder-robot/book_final_code/Chapter16/repository"
)

type PostTeamOrderService struct {
	TeamRepo *repository.TeamRepo
	Repo     *repository.TeamPostOrderRepo
}

func (t *PostTeamOrderService) PostTeamOrder(order param.TeamPostOrder) (string, error) {
	//到数据库查看是否有这个团购优惠
	teamDetail := t.TeamRepo.GetTeamDetail(order.TeamDetailId)
	if teamDetail == nil {
		return "", errors.New("TeamDetailId参数错误")
	}
	//下单数量不能小于1
	if order.Quantity < 1 {
		return "", errors.New("Quantity参数错误")
	}
	//售卖价格要大于0
	if order.RealPrice < 0 {
		return "", errors.New("RealPrice参数错误")
	}
	//下单人的手机号，不能为空
	if order.Mobile == "" {
		return "", errors.New("Mobile参数错误")
	}

	id, _ := uuid.GenerateUUID()
	o := model.TeamPostOrder{
		TeamPostOrderId: id,
		TeamDetailId:    order.TeamDetailId,
		RealPrice:       order.RealPrice,
		Quantity:        order.Quantity,
		Mobile:          order.Mobile,
		Name:            order.Name,
		Sex:             order.Sex,
		Message:         order.Message,
	}

	return t.Repo.Save(o), nil
}
