package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/hashicorp/go-uuid"
)

func ProcessTraceID() gin.HandlerFunc{
	return func(context *gin.Context) {
		traceId:= context.Request.Header.Get("X-Trace-Id")
		if traceId == "" {
			u4id, _ := uuid.GenerateUUID()
			traceId = u4id
		}
		context.Set("X-Trace-Id",traceId)
		context.Writer.Header().Set("X-Trace-Id",traceId)
		context.Next()
	}
}
