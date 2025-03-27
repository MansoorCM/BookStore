package main

import (
	"fmt"
	"net/http"
)

func orderHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Here's your order summary.")
}
