package main

import "github.com/gin-gonic/gin"

type user struct {
	ID   int
	Name string
	Age  int
}


func main() {
	r := gin.Default()
	r.GET("/hello", func(c *gin.Context) {
		c.JSON(200, gin.H{"message":"hello world"})
	})
	r.GET("/users/123", func(c *gin.Context) {
		c.JSON(200,user{ID:123,Name:"欢喜哥",Age:20})
	})
	r.Run(":8080")

}
