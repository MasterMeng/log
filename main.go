package main

import (
	"fmt"

	"github.com/monstermeng92/log/format"
)

func main() {
	log, err := format.NewLogger("./", "console") //zap.AddCaller()为显示文件名和行号，可省略
	if err != nil {
		panic(err)
	}

	log.Info("hello info")
	log.Debug("hello debug")
	log.Error("hello error")

	sugar := log.Sugar()

	fmt.Println("========== sugar ==========")

	sugar.Info("hello info")
	sugar.Debug("hello debug")
	sugar.Error("hello error")
}
