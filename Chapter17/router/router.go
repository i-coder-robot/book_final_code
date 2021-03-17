package router

import (
	"github.com/gin-gonic/gin"
	"github.com/i-coder-robot/book_final_code/Chapter17/handler"
	"net/http"
)

func Load(engine *gin.Engine, middlewares ...gin.HandlerFunc) *gin.Engine {
	engine.Use(gin.Recovery())
	engine.Use(middlewares...)
	engine.NoRoute(func(context *gin.Context) {
		context.String(http.StatusNotFound, "API路由不正确.")
	})

	account := engine.Group("/v1/account")
	{

		account.GET("", handler.AccountHandle.ListAccount)              // 获取用户列表
		account.GET("/:account_name", handler.AccountHandle.GetAccount) // 获取指定用户的详细信息
		account.POST("", handler.AccountHandle.AccountCreate)           //新增用户
		account.DELETE("/:id", handler.AccountHandle.Delete)            // 删除用户
		account.PUT("/", handler.AccountHandle.Update)                  // 更新用户
		account.POST("/login", handler.AccountHandle.Login)             //登录
		account.POST("/wxlogin", handler.AccountHandle.WXLogin)

	}
	return engine
}
