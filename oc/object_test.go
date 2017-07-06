package oc_test

import (
	"testing"
	"io/ioutil"
	"log"
	"os"
	"github.com/hongkailiu/svt-go/oc"
	"github.com/stretchr/testify/assert"
)

func TestGetResponse(t *testing.T) {
	assert := assert.New(t)
	dir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	text, err := ioutil.ReadFile(dir+"/../file_for_test/pod_response.json")
	if err != nil {
		log.Fatal(err)
	}
	t.Log("content from file" + string(text))
	response := oc.GetResponse(string(text))
	assert.Equal("v1", response.APIVersion, "they should be equal")
	if assert.NotNil(response) {
		assert.Equal("v1", response.APIVersion)
	}
}
