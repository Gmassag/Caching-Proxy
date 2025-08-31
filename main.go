package main

import (
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("Starting caching proxy...")
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Received request:", r.URL.Path)
		fmt.Fprintln(w, "Hello from caching proxy!")
	})
	http.ListenAndServe(":8080", nil)
}
