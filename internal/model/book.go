package model

import "errors"

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
