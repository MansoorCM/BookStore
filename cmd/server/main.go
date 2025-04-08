package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/MansoorCM/gobook/internal/api"
)

func main() {
	fmt.Println("welcome")

	mux := http.NewServeMux()

	handler := &api.MessageHandler{Message: "welcome to the bookstore"}

	mux.Handle("/", handler)
	mux.HandleFunc("/books", api.BooksHandler)
	mux.HandleFunc("/order", api.OrderHandler)

	fmt.Println("Starting the sever at port 8080")
	if err := http.ListenAndServe(":8080", mux); err != nil {
		log.Fatal("couldn't start http server")
	}

}
