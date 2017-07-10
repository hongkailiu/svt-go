package arg

import (
	"gopkg.in/alecthomas/kingpin.v2"
)

const AppVersion = "0.0.1-SNAPSHOT"

type AppArgs struct {
	Version string
	File    string
}

var (
	MyAppArgs = AppArgs{}
)

func init() {

}

func Parse() {
	kingpin.Version(AppVersion)
	file := kingpin.Flag("file", "This is the input config file used to define the test.").Default("conf/pyconfig.yaml").Short('f').String()
	MyAppArgs.File = *file
	kingpin.Parse()
}
