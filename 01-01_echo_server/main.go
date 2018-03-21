package main

import (
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"

	"github.com/deme0607/real-world-http-samples/constants"
)

func handler(w http.ResponseWriter, r *http.Request) {
	dump, err := httputil.DumpRequest(r, true)
	if err != nil {
		http.Error(w, fmt.Sprint(err), http.StatusInternalServerError)
		return
	}

	fmt.Println(string(dump))
	fmt.Fprint(w, "<html><body>hello</body></html>\n")
}

func main() {
	p := fmt.Sprintf(":%d", constants.Port)

	http.HandleFunc("/", handler)

	log.Println("start http listening ", p)
	log.Println(http.ListenAndServe(p, nil))
}
