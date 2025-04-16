package model

import (
	"errors"
	"slices"
)

type Book struct {
	ID     string
	Title  string
	Author string
}

var books = []*Book{
	{ID: "1", Title: "Think and Grow Rich", Author: "Nepolean Hill"},
	{ID: "2", Title: "The power of positive thinking", Author: "Norman Vincent Peale"},
}

func getBookByID(id string) (*Book, error) {
	for _, book := range books {
		if book.ID == id {
			return book, nil
		}
	}
	return nil, errors.New("Book not found")
}

func addBook(book *Book) {
	books = append(books, book)
}

func updateBook(id string, updatedBook *Book) error {
	for i, book := range books {
		if book.ID == id {
			books[i] = updatedBook
			return nil
		}
	}
	return errors.New("Book not found")
}

func deleteBook(id string) error {
	for i, book := range books {
		if book.ID == id {
			books = slices.Delete(books, i, i+1)
			return nil
		}
	}
	return errors.New("Book not found")
}
