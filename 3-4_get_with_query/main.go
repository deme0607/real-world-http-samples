package main

import (
	"net/url"
	"net/http"
	"fmt"
	"io/ioutil"
	"log"
)

const port = 18888

func main() {
	values := url.Values{
		"query": {"hello world"},
	}

	resp, err := http.Get(fmt.Sprintf("http://localhost:%d?%s", port, values.Encode()))
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
