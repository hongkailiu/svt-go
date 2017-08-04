package oc

import (
	"github.com/hongkailiu/svt-go/task"
	"github.com/hongkailiu/svt-go/log"
	"fmt"
	"sync"
	"strings"
	"errors"
	"time"
	"strconv"
	"path/filepath"
	"os"
	"io/ioutil"
)

type projectHandler struct {
	project *task.Project
	wg *sync.WaitGroup
	tuningSet *task.TuningSet
}


func (ph *projectHandler) handle() error {
	defer ph.wg.Done()
	log.Debug(fmt.Sprintf("handle project: %s", ph.project.Basename))
	for i := 0; i < ph.project.Number; i++ {
		projectName := fmt.Sprintf("%s%d", ph.project.Basename, i)
		ph.handleTemplates(projectName)

	}
	return nil
}

func HandleProject(p *task.Project, wg *sync.WaitGroup, tuningSet *task.TuningSet) {
	ph:=projectHandler{project:p, wg:wg, tuningSet:tuningSet}
	queueProjectInPool(&ph)
}

func (ph *projectHandler) handleTemplates(projectName string) error {
	stepSize := ph.tuningSet.PodsInTuningSet.Stepping.StepSize
	pause := ph.tuningSet.PodsInTuningSet.Stepping.Pause
	delay := ph.tuningSet.PodsInTuningSet.RateLimit.Delay
	for _, template := range ph.project.Templates {
		for i := 1; i <= template.Number; i++ {
			if err := handleTemplate(projectName, i-1, template); err !=nil {
				log.Critical(err)
			}
			if i % stepSize == 0 {
				sleep(pause)
			}
			sleep(delay)
		}
	}
	return nil
}


func handleTemplate(projectName string, i int, t task.Template) error {
	m := make(map[string]string)
	m["IDENTIFIER"] = strconv.Itoa(i)
	m["NAMESPACE"] = projectName
	for k, v := range t.Parameters {
		m[k]=v
	}

	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		return err
	}
	fileFullPath := fmt.Sprintf("%s%c%s", dir, os.PathSeparator, t.File)
	log.Info("fileFullPath: " + fileFullPath)

	if _, err := os.Stat(fileFullPath); os.IsNotExist(err) {
		return errors.New(fmt.Sprintf("file does not exist: %s", fileFullPath))
	}

	output, err := process(fileFullPath, m)
	if err != nil {
		return err
	}

	tmpfile, err := ioutil.TempFile(os.TempDir(), "temp")
	defer os.Remove(tmpfile.Name())
	if err != nil {
		return err
	}

	if err:=ioutil.WriteFile(tmpfile.Name(), output, 0644);err != nil {
		return err
	}

	create(tmpfile.Name())

	return nil
}

func sleep(timeS string) error {
	log.Debug("timeS:" + timeS)
	nAndUnit := strings.Split(timeS, " ")
	if len(nAndUnit) != 2 {
		msg := fmt.Sprintf("wrong time string: %s.", timeS)
		log.Critical(msg)
		return errors.New(msg)
	}

	n, err := strconv.Atoi(nAndUnit[0])
	if err != nil {
		msg := fmt.Sprintf("wrong time string: %s.", timeS)
		log.Critical(msg)
		return errors.New(msg)
	}
	switch nAndUnit[1] {
	case "s":
		time.Sleep( time.Duration(n) * time.Second)
	case "min":
		time.Sleep( time.Duration(n) * time.Minute)
	case "ms":
		time.Sleep( time.Duration(n) * time.Millisecond)
	case "hr":
		time.Sleep( time.Duration(n) * time.Hour)
	default:
		msg := fmt.Sprintf("wrong time string: %s.", timeS)
		log.Critical(msg)
		return errors.New(msg)
	}
	return nil
}