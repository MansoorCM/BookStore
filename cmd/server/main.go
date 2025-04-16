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
	r.HandleFunc("/books", api.GetBooks).Methods("GET")
	r.HandleFunc("/books/{id:[0-9]+}", api.GetBook).Methods("GET")
	r.HandleFunc("/books/{id:[0-9]+}", api.UpdateBook).Methods("PUT")
	r.HandleFunc("/books/{id:[0-9]+}", api.DeleteBook).Methods("DELETE")
	r.HandleFunc("/order", api.OrderHandler).Methods("GET")

	fmt.Println("Starting the sever at port 8080")
	if err := http.ListenAndServe(":8080", r); err != nil {
		log.Fatal("couldn't start http server")
	}

}
