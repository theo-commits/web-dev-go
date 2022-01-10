package main

import (
	"fmt"
	"net/http"
)

func handlerFunc(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	if r.URL.Path == "/" {
		fmt.Fprint(w, "<h1>Welcome to my Website!</h1>")
	} else if r.URL.Path == "/contact" {
		fmt.Fprint(w, "To get into contact with me, please send an email to <a href=\"mailto:me@tlindsey.cloud\"> me@tlindsey.cloud</a>.")
	}

}

func main() {
	http.HandleFunc("/", handlerFunc)
	http.ListenAndServe(":3000", nil)
}
