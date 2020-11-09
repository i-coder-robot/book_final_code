package main

import (
	"book-code/Chapter13/13-3/MyLog"
	"book-code/Chapter13/13-4/middleware"
	"flag"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"log"
	"net/http"
)

var (
	cfg = flag.String("config", "c", "")
)

func main() {
	flag.Parse()

	// init config
	//if err := config.Init(*cfg); err != nil {
	//	panic(err)
	//}
	//model.DB.Init()
	//defer model.DB.Close()
	r := gin.New()
	Load(
		r,
		middleware.ProcessTraceID(),
		middleware.Logging(),
	)
	port:=viper.GetString("addr")
	MyLog.Log.Info("开始监听http地址", port)
	log.Printf(http.ListenAndServe(port, r).Error())
}
