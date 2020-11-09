package log

import (
	"bytes"
	"github.com/gin-gonic/gin"
)

type BodyLoggerWriter struct {
	gin.ResponseWriter
	Body *bytes.Buffer
}

func (w BodyLoggerWriter) Write(b []byte) (int, error) {
	w.Body.Write(b)
	return w.ResponseWriter.Write(b)
}
