package main

import (
	"fmt"
	"net/http"
)

func hello(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprint(w, "<p>Hello Go Web Dev</p></p>")
}

func main() {
	fmt.Printf("Server SStarted on http://127.0.0.1:3000")
	http.HandleFunc("/", hello)
	http.ListenAndServe("127.0.0.1:3000", nil)
}
