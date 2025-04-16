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

var Books = []*Book{
	{ID: "1", Title: "Think and Grow Rich", Author: "Nepolean Hill"},
	{ID: "2", Title: "The power of positive thinking", Author: "Norman Vincent Peale"},
}

func GetBookByID(id string) (*Book, error) {
	for _, book := range Books {
		if book.ID == id {
			return book, nil
		}
	}
	return nil, errors.New("Book not found")
}

func AddBook(book *Book) {
	Books = append(Books, book)
}

func UpdateBook(id string, updatedBook *Book) error {
	for i, book := range Books {
		if book.ID == id {
			Books[i] = updatedBook
			return nil
		}
	}
	return errors.New("Book not found")
}

func DeleteBook(id string) error {
	for i, book := range Books {
		if book.ID == id {
			Books = slices.Delete(Books, i, i+1)
			return nil
		}
	}
	return errors.New("Book not found")
}
