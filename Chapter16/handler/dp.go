package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/i-coder-robot/book_final_code/Chapter16/handler/param"
	"github.com/i-coder-robot/book_final_code/Chapter16/service/dp_service"
	"io/ioutil"
	"net/http"
	"os"
)

const domain = "http://192.168.0.104:8080"

type DiscountHandler struct {
	Srv *dp_service.DiscountService
}
type RestaurantNavHandler struct {
	Srv *dp_service.RestaurantService
}

type HotelDetailHandler struct {
	Srv *dp_service.HotelService
}

type TeamDetailHandler struct {
	Srv *dp_service.TeamService
}

type OrderSeatHandler struct {
	Srv *dp_service.OrderSeatService
}

type TakeOutHandler struct {
	Srv *dp_service.TakeOutService
}

type MeHandler struct {
	Srv *dp_service.MeService
}

type SuggestFoodHandler struct {
	Srv *dp_service.SuggestFoodService
}

type SuggestHandler struct {
	Srv *dp_service.SuggestService
}

type GuessHandler struct {
	Srv *dp_service.GuessService
}
type NavHandler struct {
	Srv *dp_service.ListNavItemService
}

type PostTeamOrderHandler struct {
	Srv *dp_service.PostTeamOrderService
}

func (h *DiscountHandler) DiscountList(c *gin.Context) {
	discounts := h.Srv.ListDiscounts()
	c.JSON(http.StatusOK, gin.H{
		"items": discounts,
	})
}

func (h *RestaurantNavHandler) RestaurantNav(c *gin.Context) {
	items := h.Srv.ListRestaurantNav(1)
	c.JSON(http.StatusOK, gin.H{
		"items": items,
	})
}

func (h *RestaurantNavHandler) GoodRestaurantBillBoardHandler(c *gin.Context) {
	items := h.Srv.ListRestaurantNav(2)
	c.JSON(http.StatusOK, gin.H{
		"items": items,
	})
}

func (h *RestaurantNavHandler) GoodRestaurantTabItemHandler(c *gin.Context) {
	items := h.Srv.ListGoodRestaurantTabItem()
	c.JSON(http.StatusOK, gin.H{
		"items": items,
	})
}

func (h *HotelDetailHandler) HotelDetailHandler(c *gin.Context) {
	hotelId := c.Param("id")

	hotel, err := h.Srv.GetHotelDetailByID(hotelId)
	//TODO 直接记录错误日志
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"item": nil,
			"msg":  err.Error(),
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"item": hotel,
			"msg":  "",
		})
	}
}

func (h *TeamDetailHandler) TeamDetailHandler(c *gin.Context) {
	teamId := c.Param("id")
	detail := h.Srv.GetTeamDetail(teamId)

	c.JSON(http.StatusOK, gin.H{
		"item": detail,
	})
}

func (h *PostTeamOrderHandler) TeamOrderHandler(c *gin.Context) {
	var p param.TeamPostOrder
	err := c.BindJSON(&p)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"err": err.Error(),
		})
	}
	id, err := h.Srv.PostTeamOrder(p)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"id":  "",
			"err": err.Error(),
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"id":  id,
			"err": "",
		})
	}
}

func (h *OrderSeatHandler) OrderSeat(c *gin.Context) {
	var p param.OrderSeat
	err := c.BindJSON(&p)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"err": err.Error(),
		})
	}
	id := h.Srv.OrderSeatOp(p)
	c.JSON(http.StatusOK, gin.H{
		"id":  id,
		"err": "",
	})
}

func (h *TakeOutHandler) GetTakeOutByHotelId(c *gin.Context) {
	hotelId := c.Param("hotelId")
	item := h.Srv.GetTakeOutListByHotelId(hotelId)
	c.JSON(http.StatusOK, gin.H{
		"item": item,
	})
}

func (h *MeHandler) MeHandler(c *gin.Context) {
	items := h.Srv.ListMe()
	c.JSON(http.StatusOK, gin.H{
		"items": items,
	})
}

func (h *NavHandler) NavHandler(c *gin.Context) {
	items := h.Srv.ListNavItems(1)
	c.JSON(http.StatusOK, gin.H{
		"items": items,
	})
}

func (h *NavHandler) SubNavHandler(c *gin.Context) {
	items := h.Srv.ListNavItems(2)
	c.JSON(http.StatusOK, gin.H{
		"items": items,
	})
}

func (h *SuggestFoodHandler) TeamHandler(c *gin.Context) {
	items := h.Srv.ListSuggestList(1)
	c.JSON(http.StatusOK, gin.H{"item": items})
}

func (h *SuggestFoodHandler) RushHandler(c *gin.Context) {
	items := h.Srv.ListSuggestList(2)
	c.JSON(http.StatusOK, gin.H{"item": items})
}

func (h *SuggestHandler) TeamHandler(c *gin.Context) {
	fmt.Println("TeamHandler....")
	item := h.Srv.GetSuggestByLevel(1)
	fmt.Println(item)
	c.JSON(http.StatusOK, gin.H{"item": item})
}

func (h *SuggestHandler) RushHandler(c *gin.Context) {
	item := h.Srv.GetSuggestByLevel(2)
	c.JSON(http.StatusOK, gin.H{"item": item})
}

func (h *GuessHandler) Guess(c *gin.Context) {
	items := h.Srv.ListGuess()
	c.JSON(http.StatusOK, gin.H{
		"items": items,
	})
}

func ImageHandler(c *gin.Context) {
	imageName := c.Query("imageName")
	fmt.Println(imageName)
	dir, _ := os.Getwd()
	file, ok := ioutil.ReadFile(dir + "/static/images/" + imageName + ".png")
	if ok != nil {
		fmt.Println(ok.Error())
	}
	c.Writer.WriteString(string(file))
}
