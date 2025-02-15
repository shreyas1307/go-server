package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

func main() {

	http.HandleFunc("/hello-world", func(w http.ResponseWriter, r *http.Request) {
		fmt.Printf("Hello world")
		io.WriteString(w, "Hello world")
	})

	log.Fatal(http.ListenAndServe(":8080", nil))

}
