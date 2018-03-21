package main

import (
	"log"
	"net/http"
	"net/url"

	"github.com/deme0607/real-world-http-samples/constants"
)

func main() {
	values := url.Values{
		"test": {"value"},
	}

	resp, err := http.PostForm(constants.BaseURL, values)
	if err != nil {
		panic(err)
	}

	log.Println("Status:", resp.Status)
}
