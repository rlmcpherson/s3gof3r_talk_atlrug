package main

import (
	"fmt"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, %s.", r.URL.Path[1:])
}

func main() {
	fmt.Println("Running")
	http.ListenAndServe("127.0.0.1:8080",
		http.HandlerFunc(handler))
}
