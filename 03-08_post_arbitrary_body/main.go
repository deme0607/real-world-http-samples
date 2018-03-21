package main

import (
	"fmt"
	"go/build"
	"log"
	"net/http"
	"os"

	"github.com/deme0607/real-world-http-samples/constants"
)

const filePath = "github.com/deme0607/real-world-http-samples/03-08_post_arbitrary_body/main.go"

func main() {
	file, err := os.Open(fmt.Sprintf("%s/src/%s", build.Default.GOPATH, filePath))
	if err != nil {
		panic(err)
	}

	resp, err := http.Post(constants.BaseURL, "text/plain", file)
	if err != nil {
		panic(err)
	}

	log.Println("Status:", resp.Status)
}
