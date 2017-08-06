package http

import (
	"time"
	"github.com/hongkailiu/svt-go/log"
	"github.com/hongkailiu/svt-go/version"
	"net"
)

type info struct {
	Version string `json:"version"`
	Ips []string `json:"ips"`
	Now time.Time `json:"now"`
}

func GetInfo() *info {

	i := info{}
	i.Version = version.GetVersion()
	i.Ips = getIps()
	i.Now = time.Now()
	return &i
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