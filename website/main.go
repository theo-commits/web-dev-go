package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

// Home page for website, uses Header.().Set() func to read html tags
func home(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	fmt.Fprint(w, "<h1> Welcome to my Website! </h1>")

}

// contact page for website
func contact(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	fmt.Fprint(w, "To get into contact with me, please send an email to <a href=\"mailto:me@tlindsey.cloud\"> me@tlindsey.cloud</a>.")
}

// Create FAQ page
func faq(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	fmt.Fprint(w, "<h1>SO you have some questions?</h1> <p>You should probably just google them tbh :/</p>")
}

// Error handling for none exisitant pages. TODO: Add background image
func error404(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	fmt.Fprint(w, "<h1> Hm, that's not right... </h1> <p> Penguin_404</p>")
}

func main() {
	var e http.Handler = http.HandlerFunc(error404)
	r := mux.NewRouter()
	r.HandleFunc("/", home)
	r.HandleFunc("/contact", contact)
	r.HandleFunc("/faq", faq)
	r.NotFoundHandler = e // replace defautl 404 response with custom response
	http.ListenAndServe(":3000", r)
}
