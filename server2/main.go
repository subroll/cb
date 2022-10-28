package main

import (
	"io"
	"log"
	"net/http"
)

func main() {
	helloHandler := func(w http.ResponseWriter, req *http.Request) {
		io.WriteString(w, "Hello, server2!\n")
	}

	http.HandleFunc("/", helloHandler)
	log.Fatal(http.ListenAndServe(":3001", nil))
}
