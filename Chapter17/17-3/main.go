package main

import (
	MyLog "book-code/Chapter13/17-3/MyLog"
)

func main() {
	MyLog.Log.Info("Info日志开始")
	MyLog.Log.Error(" Eroor错误日志")
	MyLog.Log.Info("Info日志结束")
}
