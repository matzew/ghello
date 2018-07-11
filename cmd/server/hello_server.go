package main

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/matzew/ghello/pkg/config"
)

func helloWorld(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>Hello Golang</h1>")
}

func postLogger(w http.ResponseWriter, r *http.Request) {
	//	fmt.Printf("The type is: %T \n", r.Body)

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(body))
}

func main() {
	config := config.GetConfig()

	http.HandleFunc("/", helloWorld)
	http.HandleFunc("/post", postLogger)
	http.ListenAndServe(config.ListenAddress, nil)
}
