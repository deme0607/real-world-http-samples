package main

import (
	"io/ioutil"
	"log"
	"net/http"

	"github.com/deme0607/real-world-http-samples/constants"
)

const port = 18888

func main() {
	resp, err := http.Get(constants.BaseURL)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	log.Println(string(body))
	log.Println("Status:", resp.Status)
	log.Println("StatusCode:", resp.StatusCode)
	log.Println("Headers:", resp.Header)
	log.Println("Content-Length:", resp.Header.Get("Content-Length"))
}
