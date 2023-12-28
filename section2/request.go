/**
 * @author Firstname Lastname <firstname.lastname@example.com>
 * @file Description
 * @desc Created on 2023-12-28 4:49:38 am
 * @copyright APPI SASU
 */
package main

import (
	"fmt"
	"net/http"
)

func RequestHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html, charset=utf-8")
	fmt.Fprint(w, r.Method+"<br/>")
	fmt.Fprint(w, r.URL.Path+"<br/>")
	fmt.Fprint(w, r.URL.RawQuery+"<br/>")
	fmt.Fprint(w, r.Header)
	fmt.Fprint(w, r.RequestURI+"<br/>")
}

func main() {
	fmt.Printf("Server Started on http://127.0.0.1:3000")
	http.HandleFunc("/", RequestHandler)
	http.ListenAndServe("127.0.0.1:3000", nil)
}
