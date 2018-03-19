package main

import (
	"os"
	"net/http"
	"fmt"
	"log"
	"go/build"
)

const port = 18888
const filePath = "/github.com/deme0607/real-world-http-samples/3-8_post_arbitrary_body/main.go"

func main() {
	file, err := os.Open(fmt.Sprintf("%s/src/%s", build.Default.GOPATH, filePath))
	if err != nil {
		panic(err)
	}

	resp, err := http.Post(fmt.Sprintf("http://localhost:%d", port), "text/plain", file)
	if err != nil {
		panic(err)
	}

	log.Println("Status:", resp.Status)
}
