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
	result := QueueCommandInPool(fmt.Sprintf("oc get project %s", project))
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
	result := QueueCommandInPool("oc whoami")
	result.Wait()
	if err := result.Error(); err != nil {
		log.Critical(err.Error())
		return err
	}

	output := result.Value().([]byte)
	log.Info(fmt.Sprintf("%q", string(output)))
	return nil
}


func NewProject(project string) error {
	result := QueueCommandInPool(fmt.Sprintf("oc new-project %s", project))
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
	result := QueueCommandInPool(command)
	result.Wait()
	if err := result.Error(); err != nil {
		log.Critical(err.Error())
		return err
	}

	output := result.Value().([]byte)
	log.Debug(fmt.Sprintf("%q", string(output)))
	return nil
}

func process(file string, m map[string]string) ([]byte, error) {
	command := fmt.Sprintf("oc process -f %s", file)
	for k, v := range m {
		command = fmt.Sprintf("%s -p %s=%s", command, k, v)
	}
	result := QueueCommandInPool(command)
	result.Wait()
	if err := result.Error(); err != nil {
		log.Critical(err.Error())
		return nil, err
	}

	output := result.Value().([]byte)
	log.Debug(fmt.Sprintf("%q", string(output)))
	return output, nil
}


func create(file string) ([]byte, error) {
	result := QueueCommandInPool(fmt.Sprintf("oc create -f %s", file))
	result.Wait()
	if err := result.Error(); err != nil {
		log.Critical(err.Error())
		return nil, err
	}

	output := result.Value().([]byte)
	log.Debug(fmt.Sprintf("%q", string(output)))
	return output, nil
}

func isPodRunning(project string, pod string) (bool, error) {
	result := QueueCommandInPool(fmt.Sprintf("oc get pod -n %s %s", project, pod))
	result.Wait()
	if err := result.Error(); err != nil {
		log.Critical(err.Error())
		return false, err
	}

	output := string(result.Value().([]byte))
	log.Debug(fmt.Sprintf("%q", output))
	return strings.Contains(output, " Running "), nil
}