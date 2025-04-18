package api

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/MansoorCM/gobook/internal/model"
	"github.com/gorilla/mux"
)

func GetBooks(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(model.Books)
}

func GetBook(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	bookId := vars["id"]
	if bookId != "" {
		fmt.Fprintf(w, "Retrieving book with id %s", bookId)
	}
}

func CreateBook(w http.ResponseWriter, r *http.Request) {
	var book model.Book
	if err := json.NewDecoder(r.Body).Decode(&book); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	model.AddBook(&book)
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(book)
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
