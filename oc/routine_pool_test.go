package oc

import (
	"testing"
	"io/ioutil"
	"os"
	"github.com/stretchr/testify/assert"
	"fmt"
)

func TestBashPool(t *testing.T) {
	assert := assert.New(t)
	tmpfile, err := ioutil.TempFile(os.TempDir(), "temp")

	if err != nil {
		assert.Fail(err.Error())
	}

	if _, err := os.Stat(tmpfile.Name()); os.IsNotExist(err) {
		assert.Fail(fmt.Sprintf("file is not created: %s", tmpfile.Name()))
	}
	os.Remove(tmpfile.Name())

	if _, err := os.Stat(tmpfile.Name()); !os.IsNotExist(err) {
		assert.Fail(fmt.Sprintf("file is not deleted: %s", tmpfile.Name()))
	}

	StartPool(10)

	defer ClosePool()

	var command = fmt.Sprintf("touch %s", tmpfile.Name())

	result := QueueCommandInPool(command)
	result.Wait()
	if err := result.Error(); err != nil {
		assert.Fail(err.Error())
	}

	// do stuff with user
	output := result.Value().([]byte)
	assert.Empty(output)

	if _, err := os.Stat(tmpfile.Name()); os.IsNotExist(err) {
		assert.Fail(fmt.Sprintf("file is not created: %s", tmpfile.Name()))
	}

	if error := os.Remove(tmpfile.Name()); error != nil {
		assert.Fail(err.Error())
	}

}