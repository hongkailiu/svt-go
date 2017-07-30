package http

import (
	"time"
	"path/filepath"
	"os"
	"github.com/hongkailiu/svt-go/log"
	"io/ioutil"
	"net"
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
	i.Ips = getIps()
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
		log.Critical(err.Error())
		return err.Error()
	}
	return string(text)
}

func getIps() []string {
	interfaces, err := net.Interfaces()
	if err != nil {
		log.Critical(err.Error())
		return []string{err.Error()}
	}
	// handle err
	result := []string{}
	for _, i := range interfaces {
		addrs, err := i.Addrs()
		if err != nil {
			result = append(result, err.Error())
		}
		for _, addr := range addrs {
			var ip net.IP
			switch v := addr.(type) {
			case *net.IPNet:
				ip = v.IP
			case *net.IPAddr:
				ip = v.IP
			}
			result = append(result, ip.String())
		}
	}
	return result
}