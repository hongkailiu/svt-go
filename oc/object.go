package oc

import (
//"fmt"
//"github.com/hongkailiu/test-go/stringutil"
//"github.com/op/go-logging"
//"os"
)
import (
//"fmt"
	"encoding/json"
)

type Kind int

const (
	// iota is reset to 0
	pod Kind = iota  // c0 == 0
	build   // c1 == 1
)

type Response struct {
	APIVersion string `json:"apiVersion"`
	items      []Object
}

type Object struct {
	APIVersion string `json:"apiVersion"`
	Kind Kind `json:"kind"`
	metadata   struct {
			   Name      string `json:"namespace"`
			   Namespace string `json:"namespace"`
		   }
}

type Pod struct {
	Object
}

func GetResponse(responseString string) Response {

	// Get byte slice from string.
	bytes := []byte(responseString)

	// Unmarshal string into structs.
	var response Response
	json.Unmarshal(bytes, &response)
	return response
}
