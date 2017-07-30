package http

import (
	"time"
)

type Info struct {
	Version string
	Ips []string
	now time.Time
}

func GetInfo() *Info {

	info := Info{}
	info.Version = "1.2.3"
	info.Ips = []string{"192.168.31.163"}
	info.now = time.Now()
	return &info
}