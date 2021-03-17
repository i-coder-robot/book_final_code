package main

import (
	"github.com/gin-gonic/gin"
	"github.com/i-coder-robot/book_final_code/Chapter16/handler"
	"github.com/i-coder-robot/book_final_code/Chapter16/health"
	"net/http"
)

func Load(engine *gin.Engine, middlewares ...gin.HandlerFunc) *gin.Engine {
	engine.Use(gin.Recovery())
	engine.Use(middlewares...)
	engine.NoRoute(func(context *gin.Context) {
		context.String(http.StatusNotFound, "API路由不正确.")
	})
	check := engine.Group("/check")
	{
		check.GET("/health", health.Health)
	}

	dp := engine.Group("/v1/dp")
	{

		//dp.GET("/index",handler.IndexHandler)

		//首页相关
		//首页导航
		dp.GET("/nav", NavHandler.NavHandler)
		dp.GET("/subnav", NavHandler.SubNavHandler)
		//首页拼团
		dp.GET("/team", SuggestHandler.TeamHandler)
		//首页秒杀
		dp.GET("/rush", SuggestHandler.RushHandler)
		//首页猜你喜欢
		dp.GET("/guess", GuessHandler.Guess)

		//找优惠页面相关
		dp.GET("/discount", DiscountHandler.DiscountList)
		//获取图片
		dp.GET("/image", handler.ImageHandler)

		//找好店相关
		//美食排行榜
		dp.GET("/restaurantNav", RestaurantNavHandler.RestaurantNav)
		//精品榜单
		dp.GET("/restaurantBillBoard", RestaurantNavHandler.GoodRestaurantBillBoardHandler)
		//附近上榜和全域上榜
		dp.GET("/restaurantTabItem", RestaurantNavHandler.GoodRestaurantTabItemHandler)

		//获取餐馆详细信息
		dp.GET("/hotel/detail/:id", HotelDetailHandler.HotelDetailHandler)
		//获取团购详细信息
		dp.GET("/team/detail/:id", TeamDetailHandler.TeamDetailHandler)
		//团购下单
		dp.POST("/team/order", PostTeamOrderHandler.TeamOrderHandler)
		//订座下单
		dp.POST("/seat/order/:hotelId", OrderSeatHandler.OrderSeat)
		//外卖
		dp.GET("/takeout/:hotelId", TakeOutHandler.GetTakeOutByHotelId)
		//我的
		dp.GET("/me", MeHandler.MeHandler)

	}

	return engine
}
