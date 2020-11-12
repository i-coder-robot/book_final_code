package main

import (
	"errors"
	"flag"
	"github.com/gin-gonic/gin"
	"github.com/i-coder-robot/book_final_code/Chapter17/17-1/router"
	"github.com/i-coder-robot/book_final_code/Chapter17/17-2/config"
	"github.com/spf13/viper"
	"log"
	"net/http"
	"time"
)

var (
	cfg = flag.String("config", "", "")
)

func main() {
	flag.Parse()

	// init config
	if err := config.Init(*cfg); err != nil {
		panic(err)
	}
	gin.SetMode(viper.GetString("runmode"))
	g := gin.New()
	middlewares := []gin.HandlerFunc{}
	router.Load(
		g,
		middlewares...,
	)

	go func() {
		if err := checkServer(); err != nil {
			log.Fatal("自检程序发生错误...", err)
		}
		log.Print("路由成功部署.")
	}()
	port := viper.GetString("addr")
	log.Printf("开始监听http地址%s", port)
	log.Printf(http.ListenAndServe(port, g).Error())

}

func checkServer() error {
	max := viper.GetInt("max_check_count")
	for i := 0; i < max; i++ {
		//发送一个GET请求给 `/check/health`，验证服务器是否成功.
		url := viper.GetString("url") + "/check/health"
		resp, err := http.Get(url)
		if err == nil && resp.StatusCode == 200 {
			return nil
		}

		// Sleep 1 second 继续重试.
		log.Print("等待路由，1秒后重试。")
		time.Sleep(time.Second)
	}
	return errors.New("无法连接到路由.")
}
