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
	fmt.Fprint(w, "Hello Go Web Dev")
}
func main() {
	fmt.Printf("Server SStarted on 3000")
	http.HandleFunc("/", hello)
	http.ListenAndServe(":3000", nil)
}
