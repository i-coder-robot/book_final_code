package main

import (
	"flag"
	"github.com/gin-gonic/gin"
	"github.com/i-coder-robot/book_final_code/Chapter16/MyLog"
	"github.com/i-coder-robot/book_final_code/Chapter16/config"
	"github.com/i-coder-robot/book_final_code/Chapter16/middleware"
	"github.com/i-coder-robot/book_final_code/Chapter17/handler"
	"github.com/i-coder-robot/book_final_code/Chapter17/model"
	"github.com/i-coder-robot/book_final_code/Chapter17/router"
	"github.com/spf13/viper"
	"log"
	"net/http"
)

var (
	cfg            = flag.String("config", "", "")
)

func init() {
	handler.InitViper()
	handler.InitDB()
	handler.InitHandler()
}

func main() {
	flag.Parse()

	if err := config.Init(*cfg); err != nil {
		panic(err)
	}
	model.DB.Init()
	defer model.DB.Close()
	r := gin.New()
	router.Load(
		r,
		middleware.ProcessTraceID(),
		middleware.Logging(),
	)
	port := viper.GetString("addr")
	MyLog.Log.Info("开始监听http地址", port)
	log.Printf(http.ListenAndServe(port, r).Error())
}
