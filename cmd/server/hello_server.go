package main

import (
	"log"
	"net/http"
	"time"
)

func helloWorld(w http.ResponseWriter, r *http.Request) {
	// Simulate at least a bit of processing time.
	time.Sleep(100 * time.Millisecond)

	w.Header().Add("kne-foo", "yo...")
	w.Header().Add("kne-bar", "blup...")
	w.WriteHeader(http.StatusNotFound)
	//if reqBytes, err := httputil.DumpRequest(r, true); err == nil {
	log.Print("Got request")
	w.Write([]byte("<!doctype html>\n<html>\n<head>\n    <title>Error Page(tm)</title>\n</head>\n<body>\n<p>Quoth the server, 404!\n</body></html>"))

	//} else {
	//	log.Printf("Error dumping the request: %+v :: %+v", err, r)
	//}

}

func main() {

	m := http.NewServeMux()
	m.HandleFunc("/", helloWorld)
	http.ListenAndServe(":8080", m)
}
