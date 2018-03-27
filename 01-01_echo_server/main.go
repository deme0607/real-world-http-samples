package main

import (
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"
	"strconv"

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

func cookie(w http.ResponseWriter, r *http.Request) {
	dump, err := httputil.DumpRequest(r, true)
	if err != nil {
		http.Error(w, fmt.Sprint(err), http.StatusInternalServerError)
		return
	}

	fmt.Println(string(dump))

	c, err := r.Cookie("count")
	if err != nil {
		c = &http.Cookie{
			Name: "count",
			Path: "/",
		}
	}

	count, _ := strconv.Atoi(c.Value)
	count++
	c.Value = strconv.Itoa(count)

	http.SetCookie(w, c)
	fmt.Fprintf(w, "<html><body>count: %d</body></html>\n", count)
}

func main() {
	p := fmt.Sprintf(":%d", constants.Port)

	http.HandleFunc("/", handler)
	http.HandleFunc("/cookie", cookie)

	log.Println("start http listening ", p)
	log.Println(http.ListenAndServe(p, nil))
}
