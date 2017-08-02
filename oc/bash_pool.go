package oc

import (
	"gopkg.in/go-playground/pool.v3"
	"github.com/hongkailiu/svt-go/log"
)

var (
	myPool pool.Pool
)

func StartPool(number uint) {
	myPool = pool.NewLimited(number)

}

func runCommand(command string) pool.WorkFunc {

	return func(wu pool.WorkUnit) (interface{}, error) {
		log.Debug("command received: " + command)
		return RunCommand(command)
	}
}

func QueueInPool(command string) pool.WorkUnit {
	return myPool.Queue(runCommand(command))
}

func ClosePool() {
	myPool.Close()
}
