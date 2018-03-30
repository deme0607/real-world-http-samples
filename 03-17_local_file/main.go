package main

import (
	"fmt"
	"go/build"
	"log"
	"net/http"
	"net/http/httputil"
)

const filePath = "github.com/deme0607/real-world-http-samples/03-17_local_file/main.go"

func main() {
	transport := &http.Transport{}
	transport.RegisterProtocol("file", http.NewFileTransport(http.Dir("/")))

	client := http.Client{
		Transport: transport,
	}

	file := fmt.Sprintf("file://%s/src/%s", build.Default.GOPATH, filePath)
	resp, err := client.Get(file)
	if err != nil {
		panic(err)
	}

	dump, err := httputil.DumpResponse(resp, true)
	if err != nil {
		panic(err)
	}

	log.Println(string(dump))
}
