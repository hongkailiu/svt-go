package oc

import (
	"github.com/hongkailiu/svt-go/log"
	"fmt"
	"strings"
)

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
	log.Debug(fmt.Sprintf("%q", string(output)))
	return true, nil
}

func WhoAmI() error {
	result := QueueInPool("oc whoami")
	result.Wait()
	if err := result.Error(); err != nil {
		log.Critical(err.Error())
		return err
	}

	output := result.Value().([]byte)
	log.Info(fmt.Sprintf("%q", string(output)))
	return nil
}


func CreateProject(project string) error {
	result := QueueInPool(fmt.Sprintf("oc new-project %s", project))
	result.Wait()
	if err := result.Error(); err != nil {
		log.Critical(err.Error())
		return err
	}

	output := result.Value().([]byte)
	log.Debug(fmt.Sprintf("%q", string(output)))
	return nil
}

func Label(k Kind, name string, key string, value string, others string) error {
	command := fmt.Sprintf("oc label %s %s %s=%s %s", strings.ToLower(k.String()), name, key, value, others)
	result := QueueInPool(command)
	result.Wait()
	if err := result.Error(); err != nil {
		log.Critical(err.Error())
		return err
	}

	output := result.Value().([]byte)
	log.Debug(fmt.Sprintf("%q", string(output)))
	return nil
}