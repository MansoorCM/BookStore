package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	fmt.Println("welcome")

	mux := http.NewServeMux()

	handler := &MessageHandler{Message: "welcome to the bookstore"}

	mux.Handle("/", handler)
	mux.HandleFunc("/books", booksHandler)
	mux.HandleFunc("/order", orderHandler)

	fmt.Println("Starting the sever at port 8080")
	if err := http.ListenAndServe(":8080", mux); err != nil {
		log.Fatal("couldn't start http server")
	}

}
