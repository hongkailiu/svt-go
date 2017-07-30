package arg

import (
	"gopkg.in/alecthomas/kingpin.v2"
	"github.com/hongkailiu/svt-go/log"
	"github.com/hongkailiu/svt-go/http"
)

const AppVersion = "0.0.1-SNAPSHOT"

type AppArgs struct {
	Version string
	File    string
}

var (
	MyAppArgs = AppArgs{}
	httpCommand     = kingpin.Command("http", "Start http server.")
)

func init() {

}

func Parse() {
	kingpin.Version(AppVersion)
	file := kingpin.Flag("file", "This is the input config file used to define the test.").Default("conf/pyconfig.yaml").Short('f').String()
	MyAppArgs.File = *file
	//kingpin.Parse()
	switch kingpin.Parse() {
	case "http":
		log.Debug("aaa")
		http.Server{8080}.Run()
	}
}
