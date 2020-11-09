package main

import (
	"github.com/i-coder-robot/book_final_code/Chapter17/17-3/MyLog"
)

func main() {
	MyLog.Log.Info("Info日志开始")
	MyLog.Log.Error(" Eroor错误日志")
	MyLog.Log.Info("Info日志结束")
}
