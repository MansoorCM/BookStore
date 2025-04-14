package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/MansoorCM/gobook/internal/api"
	"github.com/MansoorCM/gobook/internal/middleware"
	"github.com/gorilla/mux"
)

func main() {
	fmt.Println("welcome")

	r := mux.NewRouter()
	r.Use(middleware.LoggingMiddleWare)

	handler := &api.MessageHandler{Message: "welcome to the bookstore"}

	r.Handle("/", handler)
	r.HandleFunc("/books", api.BooksHandler).Methods("GET")
	r.HandleFunc("/order", api.OrderHandler).Methods("GET")

	fmt.Println("Starting the sever at port 8080")
	if err := http.ListenAndServe(":8080", r); err != nil {
		log.Fatal("couldn't start http server")
	}

}
