package router

import (
	"github.com/gin-gonic/gin"
	"net/http"
)


func Load(engine *gin.Engine, middlewares ...gin.HandlerFunc) *gin.Engine{
	engine.Use(gin.Recovery())
	engine.Use(middlewares...)
	engine.NoRoute(func(context *gin.Context) {
		context.String(http.StatusNotFound, "API路由不正确.")
	})
	check:=engine.Group("/check")
	{
		check.GET("/health",health.Health)
	}

	return engine

}

