package oc

import (
	"github.com/hongkailiu/svt-go/log"
	"fmt"
)

var (
	oc = myOC{}
)

func init() {
	oc.PoolSize = 10
}

type OC interface {
	SetPoolSize(size uint)
	IsProjectExisting(project string) (bool, error)
	CreateProject(project string) error
	label(k Kind, name string, key string, value string) error
}

type myOC struct {
	PoolSize uint
}

func SetPoolSize(size uint) {
	oc.PoolSize = size
	StartPool(oc.PoolSize)
}

func IsProjectExisting(project string) (bool, error) {
	result := QueueInPool(fmt.Sprintf("oc get project %s", project))
	result.Wait()
	if err := result.Error(); err != nil {
		log.Critical(err.Error())
	}

	output := result.Value().([]byte)
	log.Debug(string(output))
	return false, nil
}