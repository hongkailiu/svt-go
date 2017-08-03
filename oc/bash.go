package oc

import (
	"github.com/hongkailiu/svt-go/log"
	"fmt"
	"sync"
	"strings"
	"os/exec"
	"bytes"
)

func RunCommandWithWG(cmd string, wg *sync.WaitGroup) ([]byte, error) {
	log.Debug(fmt.Sprintf("command is: %s", cmd))
	parts := strings.Fields(cmd)
	head := parts[0]
	parts = parts[1:]

	command := exec.Command(head,parts...)
	var stderr bytes.Buffer
	command.Stderr = &stderr

	out, err := command.Output()
	if err != nil {
		log.Critical(fmt.Sprintf("error occurred when executing command: %s",  stderr.String()))
	}
	if wg != nil {
		wg.Done()
	}
	return out, err
}


func RunCommand(cmd string) ([]byte, error) {
	return RunCommandWithWG(cmd, nil)
}
