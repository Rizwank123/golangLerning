package main

import (
	"fmt"
	"net/http"
)

//use independent function to handle request and response

func requestHandler(w http.ResponseWriter, r *http.Request) {
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
	fmt.Println(` █████   ███   █████          █████          █████████              ███ 
░░███   ░███  ░░███          ░░███          ███░░░░░███            ░░░  
 ░███   ░███   ░███   ██████  ░███████     ░███    ░███  ████████  ████ 
 ░███   ░███   ░███  ███░░███ ░███░░███    ░███████████ ░░███░░███░░███ 
 ░░███  █████  ███  ░███████  ░███ ░███    ░███░░░░░███  ░███ ░███ ░███ 
  ░░░█████░█████░   ░███░░░   ░███ ░███    ░███    ░███  ░███ ░███ ░███ 
    ░░███ ░░███     ░░██████  ████████     █████   █████ ░███████  █████
     ░░░   ░░░       ░░░░░░  ░░░░░░░░     ░░░░░   ░░░░░  ░███░░░  ░░░░░ 
                                                         ░███           
                                                         █████          
                                                        ░░░░░  `)

	http.ListenAndServe("localhost:8000", http.HandlerFunc(requestHandler))

}
