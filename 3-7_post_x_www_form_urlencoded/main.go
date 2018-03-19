package main

import (
	"net/url"
	"net/http"
	"fmt"
	"log"
)

const port = 18888

func main() {
	values := url.Values{
		"test": {"value"},
	}

	resp, err := http.PostForm(fmt.Sprintf("http://localhost:%d", port), values)
	if err != nil {
		panic(err)
	}

	log.Println("Status:", resp.Status)
}
