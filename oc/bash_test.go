package oc_test

import (
	"testing"
	"io/ioutil"
	"log"
	"os"
	"github.com/stretchr/testify/assert"
	"fmt"
	"github.com/hongkailiu/svt-go/oc"
	"path/filepath"
	"sync"
)

func TestRunCommand(t *testing.T) {
	dir, err := ioutil.TempDir("", "example")
	if err != nil {
		log.Fatal(err)
	}

	defer os.RemoveAll(dir)

	content := []byte("temporary file's content")
	tempFile := filepath.Join(dir, "tmpfile")
	if err := ioutil.WriteFile(tempFile, content, 0666); err != nil {
		log.Fatal(err)
	}

	wg := new(sync.WaitGroup)
	commands := []string{"cat " + tempFile}
	for _, str := range commands {
		wg.Add(1)
		go runAndCheckOutput(t, str, wg, content)
	}
	wg.Wait()
}

func TestRunCommandNotExist(t *testing.T) {
	wg := new(sync.WaitGroup)
	commands := []string{"not_exist_command"}
	for _, str := range commands {
		wg.Add(1)
		go runAndCheckError(t, str, wg, "not found")
	}
	wg.Wait()
}

func runAndCheckError(t *testing.T, str string, wg *sync.WaitGroup, keywords string) {
	assert := assert.New(t)
	_, err := oc.RunCommand(str, wg)
	if err == nil {
		log.Fatal("error should show up!")
	}
	assert.Contains(err.Error(), keywords)
}

func runAndCheckOutput(t *testing.T, str string, wg *sync.WaitGroup, expectedOutput []byte) {
	assert := assert.New(t)
	out, err := oc.RunCommand(str, wg)
	if err != nil {
		log.Fatal(fmt.Sprintf("%s", err.Error()))
	}
	assert.Equal(expectedOutput, out)
}