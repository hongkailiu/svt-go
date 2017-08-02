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

	StartPool(0)

	defer ClosePool()

	var command = fmt.Sprintf("touch %s", tmpfile.Name())

	result,error := SendWork2Pool(command)
	if error != nil {
		assert.Fail(err.Error())
	}

	br, _ := result.(BashResult)
	if br.Error != nil {
		assert.Fail(err.Error())
	}

	if _, err := os.Stat(tmpfile.Name()); os.IsNotExist(err) {
		assert.Fail(fmt.Sprintf("file is not created: %s", tmpfile.Name()))
	}

	if error := os.Remove(tmpfile.Name()); error != nil {
		assert.Fail(err.Error())
	}

}