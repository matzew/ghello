package main

import (
	"fmt"
	"net/http"

	"github.com/matzew/ghello/pkg/config"
)

func helloWorld(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>Hello Golang</h1>")
}

func main() {
	config := config.GetConfig()

	http.HandleFunc("/", helloWorld)
	http.ListenAndServe(config.ListenAddress, nil)
}
