package main

import (
	"github.com/hongkailiu/svt-go/log"
	"github.com/hongkailiu/svt-go/arg"
)

func main() {
	arg.Parse()
	log.Debug("secret")
	log.Info("info1")
	log.Notice("notice")
	log.Warning("warning")
	log.Error("err")
	log.Critical("crit000")
}
