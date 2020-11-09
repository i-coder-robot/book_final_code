package main

import "github.com/gin-gonic/gin"

func main() {
	r := gin.Default()
	r.POST("/", func(c *gin.Context) {
		wechat := c.PostForm("wechat")
		c.String(200, wechat)
	})

	r.Run(":8080")

}
