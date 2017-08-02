package oc

import (
	"runtime"
	"github.com/jeffail/tunny"
	"github.com/hongkailiu/svt-go/log"
)

var (
	pool *(tunny.WorkPool)
)

func StartPool(number int) {

	numCPUs := runtime.NumCPU()
	runtime.GOMAXPROCS(numCPUs + 1) // numCPUs hot threads + one for async tasks.

	if number == 0 {
		number = numCPUs
	}

	p, error := tunny.CreatePool(number, func(object interface{}) interface{} {
		input, _ := object.(string)
		log.Debug("input: " + input)
		return RunCommandReturnOneUnit(input)
	}).Open()

	if error != nil {
		log.Fatal(error)
	}

	pool = p

}

func SendWork2Pool(command string) (interface{}, error) {
	return (*pool).SendWork(command)
}

func ClosePool() {
	(*pool).Close()
}
