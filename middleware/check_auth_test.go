package middleware

import (
	"fmt"
	"net/http"
	"testing"
)

func TestAuth(t *testing.T) {
	var myH myHandler
	h := Auth(&myH)
	switch v:= h.(type) {
	case http.Handler:
		// do nothing
	default:
		t.Error(fmt.Sprintf("type is not http.Handler, but is %T", v))
	}
}
