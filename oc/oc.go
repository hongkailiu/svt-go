package oc

import (
	"github.com/hongkailiu/svt-go/log"
	"fmt"
)

type OC interface {
	SetPoolSize(size uint)
	WhoAmI() error
	IsProjectExisting(project string) (bool, error)
	CreateProject(project string) error
	label(k Kind, name string, key string, value string) error
}

func SetPoolSize(size uint) {
	StartPool(size)
}

func IsProjectExisting(project string) (bool, error) {
	result := QueueInPool(fmt.Sprintf("oc get project %s", project))
	result.Wait()
	if err := result.Error(); err != nil {
		log.Critical(err.Error())
		return false, err
	}

	output := result.Value().([]byte)
	log.Debug(string(output))
	// TODO
	return false, nil
}

func WhoAmI() error {
	result := QueueInPool("oc whoami")
	result.Wait()
	if err := result.Error(); err != nil {
		log.Critical(err.Error())
		return err
	}

	output := result.Value().([]byte)
	log.Info( fmt.Sprintf("%q", string(output)))
	return nil
}