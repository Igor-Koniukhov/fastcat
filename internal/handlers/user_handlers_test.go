package handlers

import (
	"github.com/igor-koniukhov/fastcat/internal/models"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
)

type postData struct {
	User models.User
}

var theTests = []struct {
	name               string
	url                string
	method             string
	params             []postData
	expectedStatusCode int
}{
	{"login", "/login", "GET", []postData{}, http.StatusUnauthorized},
	{"show-login", "/show-login", "GET", []postData{}, http.StatusOK},
	{"refresh", "/refresh", "GET", []postData{}, http.StatusBadRequest},
	{"registration", "/registration", "POST", []postData{}, http.StatusOK},
}

func TestHandlers(t *testing.T) {
	routes := getRoutes()
	ts := httptest.NewServer(routes)
	defer ts.Close()
	for _, e := range theTests {
		if e.method == "GET" {
			resp, err := ts.Client().Get(ts.URL + e.url)
			if err != nil {
				t.Log(err)
				t.Fatal(err)
			}

			if resp.StatusCode != e.expectedStatusCode {
				t.Errorf("for %s expected %d but got %d", e.name, e.expectedStatusCode, resp.StatusCode)
			}
		} else {
			values := url.Values{}
			for _, x := range e.params {
				values.Add(x.User.Name, x.User.Email)
			}

			resp, err := ts.Client().PostForm(ts.URL+e.url, values)
			if err != nil {
				t.Log(err)
				t.Fatal(err)
			}

			if resp.StatusCode != e.expectedStatusCode {
				t.Errorf("for %s expected %d but got %d", e.name, e.expectedStatusCode, resp.StatusCode)
			}
		}
	}
}
