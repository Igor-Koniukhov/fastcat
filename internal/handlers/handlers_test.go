package handlers

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
)

type postData struct {
	Key string
	Value string
}

var theTests = []struct {
	name               string
	url                string
	method             string
	params             []postData
	expectedStatusCode int
}{
	{"login", "/login", "POST", []postData{}, http.StatusUnauthorized},
	{"show-login", "/show-login", "GET", []postData{}, http.StatusOK},
	{"refreshToken", "/refreshToken", "GET", []postData{}, http.StatusBadRequest},
	{"logout", "/logout", "GET", []postData{}, http.StatusOK},
	{"registration", "/registration", "POST", []postData{}, http.StatusOK},
	{"users", "/users", "GET", []postData{}, http.StatusOK},
	{"user", "/user/2", "GET", []postData{}, http.StatusOK},
	{"userCreate", "/user/create", "POST", []postData{}, http.StatusCreated},
	{"userUpdate", "/user/update/3", "PUT", []postData{}, http.StatusOK},
	{"userDelete", "/user/delete/5", "DELETE", []postData{}, http.StatusAccepted},

	{"order-create", "/order-create", "POST", []postData{}, http.StatusCreated},
	{"order", "/order/", "GET", []postData{}, http.StatusOK},
	{"orders", "/orders", "GET", []postData{}, http.StatusOK},
	{"orderUpdate", "/order/update/", "GET", []postData{}, http.StatusOK},
	{"orderDelete", "/order/delete", "DELETE", []postData{}, http.StatusAccepted},

	{"supplierCreate", "/supplier/create", "POST", []postData{}, http.StatusCreated},
	{"supplier", "/supplier/", "GET", []postData{}, http.StatusOK},
	{"suppliers", "/suppliers", "GET", []postData{}, http.StatusOK},
	{"supplierUpdate", "/supplier/update/", "PUT", []postData{}, http.StatusOK},
	{"supplierDelete", "/supplier/delete/", "DELETE", []postData{}, http.StatusAccepted},

	{"productCreate", "/product/create", "POST", []postData{}, http.StatusCreated},
	{"product", "/product/", "GET", []postData{}, http.StatusOK},
	{"products", "/products", "GET", []postData{}, http.StatusOK},
	{"productUpdate", "/product/update/", "GET", []postData{}, http.StatusOK},
	{"productDelete", "/product/delete", "DELETE", []postData{}, http.StatusAccepted},

	{"cartCreate", "/cart/create", "POST", []postData{}, http.StatusCreated},
	{"cart", "/cart/", "GET", []postData{}, http.StatusOK},
	{"carts", "/carts", "GET", []postData{}, http.StatusOK},
	{"cartUpdate", "/cart/update/", "PUT", []postData{}, http.StatusOK},
	{"cartDelete", "/cart/delete/", "DELETE", []postData{}, http.StatusAccepted},

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
				values.Add(x.Key, x.Value)
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

var theTestsAuth = []struct {
	name               string
	url                string
	method             string
	params             []postData
	expectedStatusCode int
}{

	{"user", "/user/", "GET", []postData{}, http.StatusOK},
	{"userUpdate", "/user/update/", "PUT", []postData{}, http.StatusOK},
	{"userDelete", "/user/delete/", "DELETE", []postData{}, http.StatusOK},

	{"order-create", "/order-create", "POST", []postData{}, http.StatusOK},
	{"order", "/order/", "GET", []postData{}, http.StatusOK},
	{"orders", "/orders", "GET", []postData{}, http.StatusOK},
	{"orderUpdate", "/order/update/", "GET", []postData{}, http.StatusOK},
	{"orderDelete", "/order/delete", "DELETE", []postData{}, http.StatusOK},

	{"supplierCreate", "/supplier/create", "POST", []postData{}, http.StatusOK},
	{"supplierUpdate", "/supplier/update/", "PUT", []postData{}, http.StatusOK},
	{"supplierDelete", "/supplier/delete/", "DELETE", []postData{}, http.StatusAccepted},

	{"productCreate", "/product/create", "POST", []postData{}, http.StatusOK},
	{"productUpdate", "/product/update/", "GET", []postData{}, http.StatusOK},
	{"productDelete", "/product/delete", "DELETE", []postData{}, http.StatusOK},


	{"cartUpdate", "/cart/update/", "PUT", []postData{}, http.StatusOK},
	{"cartDelete", "/cart/delete/", "DELETE", []postData{}, http.StatusOK},

}


func TestAuthHandlers(t *testing.T) {
	routes := getAuthRoutes()
	ts := httptest.NewServer(routes)
	defer ts.Close()
	for _, e := range theTestsAuth {
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
				values.Add(x.Key, x.Value)
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
