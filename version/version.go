package version

import (
	"path/filepath"
	"os"
	"io/ioutil"

	"github.com/hongkailiu/svt-go/log"
)

const (
	VersionFile = "conf/version"
)

func GetVersion() string {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		log.Critical(err)
		return err.Error()
	}

	text, err := ioutil.ReadFile(dir + string(os.PathSeparator) + VersionFile)
	if err != nil {
		log.Critical(err.Error())
		return err.Error()
	}
	return string(text)
}
