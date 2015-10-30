package main

import (
	"fmt"
	"helloworld"
	"net/http"
)

func main() {
	http.HandleFunc("/", index)

	http.ListenAndServe(":8080", nil)
}

func index(w http.ResponseWriter, r *http.Request) {
	name := r.FormValue("name")
	if name == "" {
		name = "Unknown"
	}

	fmt.Fprintln(w, helloworld.Hello(name))
}
