package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"

	"github.com/deme0607/real-world-http-samples/constants"
)

func main() {
	values := url.Values{
		"query": {"hello world"},
	}

	resp, err := http.Get(fmt.Sprintf("%s?%s", constants.BaseURL, values.Encode()))
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	log.Println(string(body))
}
