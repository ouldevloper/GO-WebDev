/**
 * @author Firstname Lastname <firstname.lastname@example.com>
 * @file Description
 * @desc Created on 2023-12-28 5:07:03 am
 * @copyright APPI SASU
 */
package main

import (
	"fmt"
	"net/http"
)

func boot(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	switch r.URL.Path {
	case "/":
		fmt.Fprint(w, "Home Page ....")
	case "/email":
		fmt.Fprint(w, "This is the Email Page")
	default:
		fmt.Fprintf(w, "Page Not Found....")
	}

}

func main() {
	fmt.Printf("Server SStarted on http://127.0.0.1:3000")
	http.HandleFunc("/", boot)
	http.ListenAndServe("127.0.0.1:3000", nil)
}
