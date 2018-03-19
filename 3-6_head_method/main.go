package main

import (
	"log"
	"net/http"

	"github.com/deme0607/real-world-http-samples/constants"
)

const port = 18888

func main() {
	resp, err := http.Head(constants.BaseURL)
	if err != nil {
		panic(err)
	}

	log.Println("Status:", resp.Status)
	log.Println("Headers:", resp.Header)
}
