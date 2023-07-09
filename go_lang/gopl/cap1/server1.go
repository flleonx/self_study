package main

import (
	"log"
	"net/http"
)

func server1() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func handler(w http.ResponseWriter, _ *http.Request) {
	lissajous(w)
	// fmt.Fprintf(w, "URL.Path = %q\n", r.URL.Path)
}
