//https://groups.google.com/forum/#!topic/golang-nuts/v1TXLIRZjv4
package http
//package http_test

import (
	"testing"
	"net/http"
	"net/http/httptest"
	"github.com/stretchr/testify/assert"
	"encoding/json"
)

func TestRootHandler(t *testing.T) {

	assert := assert.New(t)

	req, err := http.NewRequest("GET", "/", nil)
	assert.Nil(err)

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(rootHandler)

	handler.ServeHTTP(rr, req)

	assert.Equal(http.StatusOK, rr.Code)
	var i info
	json.Unmarshal(rr.Body.Bytes(), &i)
	assert.NotNil(i.Now)

	assert.Contains(rr.Body.String(), "version")
}