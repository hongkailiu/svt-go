package oc

import (
	"strings"
	"errors"
	"fmt"
)

type Kind int

const (
	// iota is reset to 0
	pod Kind = iota
	build
	unknown
)

func (kind Kind) String() string {

	switch kind {
	case pod:
		return "Pod"

	case build:
		return "Build"
	}

	return "Unknown"
}

func (k *Kind) UnmarshalJSON(b []byte) error {
	str := strings.Trim(string(b), `"`)

	switch str {
	case "Pod":
		*k = pod
	case "Build":
		*k = build
	default:
		return errors.New(fmt.Sprintf("unsuppoted kind: %s", str))
	}
	return nil
}

type Response struct {
	APIVersion string `json:"apiVersion"`
	Items      []Object `json:"items"`
}

type Object struct {
	APIVersion string `json:"apiVersion"`
	Kind       Kind `json:"kind"`
	Metadata   struct {
			   Name      string `json:"namespace"`
			   Namespace string `json:"namespace"`
		   }
}

type Pod struct {
	Object
}
