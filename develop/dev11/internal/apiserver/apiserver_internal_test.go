package apiserver

import (
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestApiServer_HandleHello(t *testing.T) {
	s := New(NewConfig("./configs/server.yaml"))
	rec := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/hello", nil)
	s.handleHello(rec, req)
	assert.Equal(t, rec.Body.String(), "Hello World")
}
