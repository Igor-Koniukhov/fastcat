package render

import (
	"github.com/igor-koniukhov/fastcat/internal/config"
	"net/http"
	"os"
	"testing"
)


var testApp config.AppConfig

func TestMain(m *testing.M) {
	app = &testApp
	os.Exit(m.Run())
}

type myWriter struct{}

func (tw *myWriter) Header() http.Header {
	var h http.Header
	return h
}

func (tw *myWriter) WriteHeader(i int) {}

func (tw *myWriter) Write(b []byte) (int, error) {
	length := len(b)
	return length, nil
}