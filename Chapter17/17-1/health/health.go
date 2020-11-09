package health

import (
	"github.com/gin-gonic/gin"
	"net/http"
)


func Health(c *gin.Context) {
	message := "OK"
	c.String(http.StatusOK, "\n"+message)
}
