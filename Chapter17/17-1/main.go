package main

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/i-coder-robot/book_final_code/Chapter17/17-1/router"
	"log"
	"net/http"
	"time"
)

func main() {
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

	log.Printf("开始监听http地址: %s", "9090")
	log.Printf(http.ListenAndServe(":9090", g).Error())
}

func checkServer() error {
	for i := 0; i < 10; i++ {
		//发送一个GET请求给 `/check/health`，验证服务器是否成功.
		resp, err := http.Get("http://127.0.0.1:9090/check/health")
		if err == nil && resp.StatusCode == 200 {
			return nil
		}

		// Sleep 1 second 继续重试.
		log.Print("等待路由，1秒后重试。")
		time.Sleep(time.Second)
	}
	return errors.New("无法连接到路由.")
}
