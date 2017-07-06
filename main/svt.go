package main

import (
	//"fmt"
	//"github.com/hongkailiu/test-go/stringutil"
	"github.com/op/go-logging"
	"os"
	"github.com/hongkailiu/svt-go/oc"
)

var log = logging.MustGetLogger("svt")

// Example format string. Everything except the message has a custom color
// which is dependent on the log level. Many fields have a custom output
// formatting too, eg. the time returns the hour down to the milli second.
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


	//fmt.Printf("hello, world\n")
	//fmt.Printf(stringutil.Reverse("!oG ,olleH"))

	response :=oc.GetResponse("{\"apiVersion\": \"v1\"}");
	log.Critical(response.APIVersion)
}
