package oc

import (
	"github.com/hongkailiu/svt-go/log"
	"fmt"
	"sync"
	"strings"
	"os/exec"
)

func RunCommand(cmd string, wg *sync.WaitGroup) ([]byte, error) {
	log.Debug(fmt.Sprintf("command is %s", cmd))
	parts := strings.Fields(cmd)
	head := parts[0]
	parts = parts[1:]
	out, err := exec.Command(head,parts...).Output()
	wg.Done()
	return out, err
}
