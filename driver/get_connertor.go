package driver

import (
	"io/ioutil"
	"log"
	"net/http"
)

func GetBodyConnection(url string) (response []byte) {
	conn, err := http.Get(url)
	if err != nil {
		log.Println(err)
	}
	defer conn.Body.Close()
	response, err = ioutil.ReadAll(conn.Body)
	if err != nil {
		log.Println(err)
	}
	return
}
