/**
 * @author Firstname Lastname <firstname.lastname@example.com>
 * @file Description
 * @desc Created on 2023-12-28 3:47:03 am
 * @copyright APPI SASU
 */
package main

import (
	"fmt"
	"net/http"
)

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "<p>Hello Go Web Dev</p></p>")
	fmt.Fprint(w, "<p>Hello Go Web Dev</p></p>")
	fmt.Fprint(w, "<p>Hello Go Web Dev</p></p>")
	fmt.Fprint(w, "<p>Hello Go Web Dev</p></p>")
}
func main() {
	fmt.Printf("Server SStarted on http://127.0.0.1:3000")
	http.HandleFunc("/", hello)
	http.ListenAndServe("127.0.0.1:3000", nil)
}
