package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	fmt.Println("welcome")

	handler := &MessageHandler{Message: "welcome to the bookstore"}

	http.Handle("/", handler)
	fmt.Println("Starting the sever at port 8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal("couldn't start http server")
	}

}
