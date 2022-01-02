package main

import (
	"fmt"
	"net/http"
)

func main() {

	http.HandleFunc("/", func(res http.ResponseWriter, req *http.Request) {
		fmt.Fprintf(res, "Sample Go HTTP server. Please use /helloworld route or /static/sample.html")
	})

	http.HandleFunc("/helloworld", func(res http.ResponseWriter, req *http.Request) {
		fmt.Fprintf(res, "<html><head><title>Hello World by GO</title></head>")
		fmt.Fprintf(res, "<body>")
		fmt.Fprintf(res, "<h1>Hello World by GO</h1>")
		fmt.Fprintf(res, "</body></html>")
	})

	fs := http.FileServer(http.Dir("static/"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	http.ListenAndServe(":3000", nil)
}
