package apiserver

import (
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestAPIServer_HandleHello(t *testing.T) {
	api := New(NewConfig())
	rec := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, "/hello", nil)

	api.handleHello().ServeHTTP(rec, req)
	assert.Equal(t, rec.Body.String(), "Hello")
}
