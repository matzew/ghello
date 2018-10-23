package main

import (
	"fmt"
	"net/http"

	"github.com/cloudevents/sdk-go/v01"
	"github.com/matzew/ghello/pkg/config"
)

func helloWorld(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>Hello Golang</h1>")
}

func postLogger(w http.ResponseWriter, r *http.Request) {

	marshaller := v01.NewDefaultHTTPMarshaller()
	// req is *http.Request
	event, err := marshaller.FromRequest(r)
	if err != nil {
		panic(err)
		//panic("Unable to parse event from http Request: " + err.String())
	}

	val, ok := event.Get("eventType")

	fmt.Printf("eventType: %s \n", val)
	if ok == false {
		panic("should have eventtype...")
	}
}

func main() {
	config := config.GetConfig()

	http.HandleFunc("/", helloWorld)
	http.HandleFunc("/post", postLogger)
	http.ListenAndServe(config.ListenAddress, nil)
}
