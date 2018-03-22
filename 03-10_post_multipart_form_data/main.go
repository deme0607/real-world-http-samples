package main

import (
	"bytes"
	"fmt"
	"go/build"
	"io"
	"log"
	"mime/multipart"
	"net/http"
	"os"

	"github.com/deme0607/real-world-http-samples/constants"
)

const filePath = "github.com/deme0607/real-world-http-samples/03-10_post_multipart_form_data/photo.jpg"

func main() {
	var buffer bytes.Buffer

	writer := multipart.NewWriter(&buffer)
	writer.WriteField("name", "Naoki Shimizu")

	fileWriter, err := writer.CreateFormFile("thumbnail", "photo.jpg")
	if err != nil {
		panic(err)
	}

	fileName := fmt.Sprintf("%s/src/%s", build.Default.GOPATH, filePath)
	readFile, err := os.Open(fileName)
	if err != nil {
		panic(err)
	}

	io.Copy(fileWriter, readFile)
	writer.Close()

	resp, err := http.Post(constants.BaseURL, writer.FormDataContentType(), &buffer)
	if err != nil {
		panic(err)
	}

	log.Println("Status:", resp.Status)
}
