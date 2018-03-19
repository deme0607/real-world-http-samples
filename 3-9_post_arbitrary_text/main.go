package main

import (
	"strings"
	"net/http"
	"github.com/deme0607/real-world-http-samples/constants"
	"log"
)

func main() {
	reader := strings.NewReader("Text")
	resp, err := http.Post(constants.BaseURL, "test/plain", reader)
	if err != nil {
		panic(err)
	}

	log.Println("Status:", resp.Status)
}
