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
	"regexp"
)

type projectHandler struct {
	projectName string
	templates *[]task.Template
	wg *sync.WaitGroup
	tuningSet *task.TuningSet
}


func (ph *projectHandler) handle() error {
	defer ph.wg.Done()
	log.Debug(fmt.Sprintf("handle project: %s", ph.projectName))
	ph.handleTemplates(ph.projectName)
	return nil
}

func HandleProject(projectName string, templates *[]task.Template, wg *sync.WaitGroup, tuningSet *task.TuningSet) {
	ph:=projectHandler{projectName:projectName, templates:templates, wg:wg, tuningSet:tuningSet}
	queueProjectInPool(&ph)
}

func (ph *projectHandler) handleTemplates(projectName string) error {
	stepSize := ph.tuningSet.PodsInTuningSet.Stepping.StepSize
	pause := ph.tuningSet.PodsInTuningSet.Stepping.Pause
	delay := ph.tuningSet.PodsInTuningSet.RateLimit.Delay
	podNames := []string{}
	for _, template := range *(ph.templates) {
		for i := 1; i <= template.Number; i++ {
			if err := handleTemplate(projectName, i-1, template, &podNames); err !=nil {
				log.Critical(err)
			}
			if i % stepSize == 0 {
				sleep(pause)
			}
			sleep(delay)
		}
	}
	checkPods(projectName, podNames)
	return nil
}
func checkPods(projectName string, podNames []string) {
	log.Debug(fmt.Sprintf("%q", podNames))
	notRunningPodNames := []string{}
	if len(podNames) > 0 {
		for _, podName := range podNames {
			running, err := isPodRunning(projectName, podName)
			if err != nil {
				log.Critical(err)
				return
			}
			if !running {
				notRunningPodNames = append(notRunningPodNames, podName)
			}
		}
		checkPods(projectName, notRunningPodNames)
	}
}


func handleTemplate(projectName string, i int, t task.Template, podNames *([]string)) error {
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

	output, err = create(tmpfile.Name())
	if err != nil {
		return err
	}
	outputStr := string(output)
	lines := strings.Split(outputStr,"\n")
	for _, line := range lines {
		if podName:=getPodName(line); podName!="" {
			*podNames = append(*podNames, podName)
		}
	}

	return nil
}

func getPodName(line string) string {
	if strings.HasPrefix(line, "pod") {
		re := regexp.MustCompile("\".*\"")
		result := re.FindString(line)
		return result[1 : len(result)-1]
	}
	return ""
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