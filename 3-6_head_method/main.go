package main

import (
	"net/http"
	"fmt"
	"log"
)

const port = 18888

func main() {
	resp, err := http.Head(fmt.Sprintf("http://localhost:%d", port))
	if err != nil {
		panic(err)
	}

	log.Println("Status:", resp.Status)
	log.Println("Headers:", resp.Header)
}
