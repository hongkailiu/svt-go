package oc

import (
	"github.com/hongkailiu/svt-go/log"
	"fmt"
	"sync"
	"strings"
	"os/exec"
)

func RunCommandWithWG(cmd string, wg *sync.WaitGroup) ([]byte, error) {
	log.Debug(fmt.Sprintf("command is %s", cmd))
	parts := strings.Fields(cmd)
	head := parts[0]
	parts = parts[1:]
	out, err := exec.Command(head,parts...).Output()
	if wg != nil {
		wg.Done()
	}
	return out, err
}


func RunCommand(cmd string) ([]byte, error) {
	return RunCommandWithWG(cmd, nil)
}
