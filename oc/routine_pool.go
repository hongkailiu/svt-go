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

func handleProject(ph *projectHandler) pool.WorkFunc {

	return func(wu pool.WorkUnit) (interface{}, error) {
		log.Debug("Basename received: " + ph.project.Basename)
		return nil, ph.handle()
	}
}

func QueueCommandInPool(command string) pool.WorkUnit {
	return myPool.Queue(runCommand(command))
}

func queueProjectInPool(ph *projectHandler) pool.WorkUnit {
	return myPool.Queue(handleProject(ph))
}

func ClosePool() {
	myPool.Close()
}
