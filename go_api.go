package main

import (
	"fmt"
	"net/http"
)

func main() {
	port := 3000
	fmt.Println("Starting Server on port", port)

	http.HandleFunc("/hello_world", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello, World"))
	})
	http.ListenAndServe(":3000", nil)
}
