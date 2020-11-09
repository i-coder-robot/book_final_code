package main

import (
	"book-code/Chapter13/13-4/handler"
	"book-code/Chapter13/13-4/health"
	"github.com/gin-gonic/gin"
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

	account := engine.Group("/v1/account")
	{
		account.POST("", AccountHandler.AccountCreate)           //新增用户
		account.GET("", AccountHandler.ListAccount)              // 获取用户列表
		account.GET("/:account_name", AccountHandler.GetAccount) // 获取指定用户的详细信息
		account.DELETE("/:id", AccountHandler.Delete)            // 删除用户
		account.PUT("/:id", AccountHandler.Update)               // 更新用户
		account.POST("/login", AccountHandler.Login)
		account.POST("/wxlogin", AccountHandler.WXLogin)

	}

	dp := engine.Group("/v1/dp")
	{

		//dp.GET("/index",handler.IndexHandler)

		//首页相关
		//首页导航
		dp.GET("/nav", NavHandler.NavHandler)
		dp.GET("/subnav", NavHandler.SubNavHandler)
		//首页拼团
		dp.GET("/team", SuggestFoodHandler.TeamHandler)
		//首页秒杀
		dp.GET("/rush", SuggestFoodHandler.RushHandler)
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
		dp.GET("/team/detail/:id",TeamDetailHandler.TeamDetailHandler)
		//团购下单
		dp.POST("/team/order/:id",PostTeamOrderHandler.TeamOrderHandler)
		//订座下单
		dp.POST("/seat/order/:hotelId",OrderSeatHandler.OrderSeat)
		//外卖
		dp.GET("/takeout/:hotelId",TakeOutHandler.GetTakeOutByHotelId)
		//我的
		dp.GET("/me", MeHandler.Me)

	}

	return engine
}