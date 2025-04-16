package api

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func GetBooks(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Retrieving all books.")
}

func GetBook(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	bookId := vars["id"]
	if bookId != "" {
		fmt.Fprintf(w, "Retrieving book with id %s", bookId)
	}
}

func UpdateBook(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	bookId := vars["id"]
	if bookId != "" {
		fmt.Fprintf(w, "Updating book with id %s", bookId)
	}
}

func DeleteBook(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	bookId := vars["id"]
	if bookId != "" {
		fmt.Fprintf(w, "Deleting book with id %s", bookId)
	}
}
