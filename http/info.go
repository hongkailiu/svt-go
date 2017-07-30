package http

import (
	"time"
	"path/filepath"
	"os"
	"github.com/hongkailiu/svt-go/log"
	"io/ioutil"
)

const (
	VersionFile = "conf/version"
)

type info struct {
	Version string `json:"version"`
	Ips []string `json:"ips"`
	Now time.Time `json:"now"`
}

func GetInfo() *info {

	i := info{}
	i.Version = getVersion()
	i.Ips = []string{"192.168.31.163"}
	i.Now = time.Now()
	return &i
}

func getVersion() string {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		log.Critical(err)
		return err.Error()
	}

	text, err := ioutil.ReadFile(dir + string(os.PathSeparator) + VersionFile)
	if err != nil {
		log.Critical("aaa" + err.Error())
		return err.Error()
	}
	return string(text)
}