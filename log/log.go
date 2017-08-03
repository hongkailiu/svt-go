package log

import (
	"github.com/op/go-logging"
	"os"
)

var log = logging.MustGetLogger("svt")

var format = logging.MustStringFormatter(
	`%{color}%{time:2006-01-02T15:04:05.000-0700} %{shortfunc} ▶ %{level:.4s} %{id:03x}%{color:reset} %{message}`,
)

func init() {
	backend := logging.NewLogBackend(os.Stdout, "", 0)
	backendFormatter := logging.NewBackendFormatter(backend, format)
	logging.SetLevel(logging.DEBUG, "")
	logging.SetBackend(backendFormatter)
}

func Debug(args ...interface{}) {
	log.Debug(args)
}

func Info(args ...interface{}) {
	log.Info(args)
}

func Notice(args ...interface{}) {
	log.Notice(args)
}

func Warning(args ...interface{}) {
	log.Warning(args)
}

func Error(args ...interface{}) {
	log.Error(args)
}

func Critical(args ...interface{}) {
	log.Critical(args)
}

func Fatal(args ...interface{}) {
	log.Fatal(args)
}