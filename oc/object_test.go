package oc_test

import (
	"testing"
	"io/ioutil"
	"log"
	"os"
	"github.com/hongkailiu/svt-go/oc"
	"github.com/stretchr/testify/assert"
	"encoding/json"
	"fmt"
)

func TestUnmarshalJSON(t *testing.T) {
	assert := assert.New(t)
	dir, err := os.Getwd()
	if err != nil {
		log.Fatal(err.Error())
	}
	text, err := ioutil.ReadFile(dir+"/../file_for_test/pod_response.json")
	if err != nil {
		log.Fatal(err.Error())
	}
	var response oc.Response
	if err2 := json.Unmarshal([]byte(text), &response); err2==nil {
		assert.Equal("v1", response.APIVersion, "they should be equal")
		if assert.NotNil(response) {
			assert.Equal("v1", response.APIVersion)
			if assert.NotNil(response.Items) {
				assert.NotEmpty(response.Items)
				assert.NotEmpty(1, len(response.Items))
				if assert.NotNil(response.Items[0]) {
					assert.Equal("v1", response.Items[0].APIVersion)
					assert.Equal("Pod", response.Items[0].Kind.String())
					assert.Equal(oc.Kind(0), response.Items[0].Kind)
				}
			}
		}
	} else {
		log.Fatal(err2.Error())
	}

}


