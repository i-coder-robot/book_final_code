package health

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// Health 输出 `OK` 作为可以访问的结果.
func Health(c *gin.Context) {
	message := "OK"
	c.String(http.StatusOK, "\n"+message)
}
