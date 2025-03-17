package main

import (
	"fmt"
	"net/http"
)

func booksHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "The bookstore site is under progress. Once done, you can see the available list of books here.")
}
