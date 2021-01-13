package main

import "github.com/gin-gonic/gin"

func main() {
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, c.QueryArray("media"))
	})
	r.GET("/map", func(c *gin.Context) {
		c.JSON(200,c.Query("ids"))
	})
	r.Run(":8080")

}
