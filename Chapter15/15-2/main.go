package main

import "github.com/gin-gonic/gin"

func main() {
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"Blog":"www.coolpest8.com",
			"wechat":"zb13161658867",
		})
	})
	r.Run(":8080")

}
