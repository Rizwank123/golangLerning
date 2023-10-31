package main

import (
	"fmt"
	"net/http"
	"strings"
)

// net/http package

type MyWebServer bool

func (m MyWebServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Welcome to new web api")
	fmt.Fprint(w, "Request is : %+v", r)
}
func displayBanner(message string) {
	border := strings.Repeat("*", len(message)+4)
	fmt.Println(border)
	fmt.Printf("* %s *\n", message)
	fmt.Println(border)
}

func main() {
	displayBanner("Welcome to new Api")
	var k MyWebServer
	http.ListenAndServe("localhost:8000", k)

}
