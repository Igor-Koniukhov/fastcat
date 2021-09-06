package driver

import (
	web "github.com/igor-koniukhov/webLogger/v3"
	"io/ioutil"
	"net/http"
)

func GetBodyConnection(url string) (response []byte) {
	conn, err := http.Get(url)
	web.Log.Error(err)
	defer conn.Body.Close()
	response, err = ioutil.ReadAll(conn.Body)
	web.Log.Error(err)
	return
}
