package main

import (
	"fmt"
	"net/http"
)

type MyWebServer bool

func (m MyWebServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, `
	<html>
	<head>
	Welcome
	</head>
	<body>
	<h1>To my Web Server</h1>
	</body>
	</htm>`)
}
func main() {
	var k MyWebServer
	http.ListenAndServe("localhost:8000", k)
}
