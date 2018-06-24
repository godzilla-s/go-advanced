package main

import (
	"fmt"
	"net/http"
)

var s = `https://github.com/hyperledger/fabric/tree/v1.0.0-preview/examples/sfhackfest`

func Handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, s)
}

func main() {
	http.HandleFunc("/", Handler)
	http.ListenAndServe(":8001", nil)
}
