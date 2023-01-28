package main

import (
	"fmt"
	"net/http"

	"github.com/NYTimes/gziphandler"
)

func main() {
	helloWorldFunc := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, %s!", r.URL.Path[1:])
	})
	compressedHelloWorldFunc := gziphandler.GzipHandler(helloWorldFunc)
	http.Handle("/hello_world", compressedHelloWorldFunc)
}
