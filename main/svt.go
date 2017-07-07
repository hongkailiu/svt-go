package main

import (
	"github.com/op/go-logging"
	"os"
)

var log = logging.MustGetLogger("svt")

var format = logging.MustStringFormatter(
	`%{color}%{time:15:04:05.000} %{shortfunc} ▶ %{level:.4s} %{id:03x}%{color:reset} %{message}`,
)

func main() {
	backend := logging.NewLogBackend(os.Stderr, "", 0)
	backendFormatter := logging.NewBackendFormatter(backend, format)
	//backend1Leveled := logging.AddModuleLevel(backend1)
	//backend1Leveled.SetLevel(logging.ERROR, "")
	//logging.SetBackend(backend1Leveled, backend1Formatter)
	logging.SetBackend(backendFormatter)

	log.Debug("secret")
	log.Info("info1")
	log.Notice("notice")
	log.Warning("warning")
	log.Error("err")
	log.Critical("crit")
}
