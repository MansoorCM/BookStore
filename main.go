package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	fmt.Println("welcome")

	http.HandleFunc("/", helloHandler)
	fmt.Println("Starting the sever at port 8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal("couldn't start http server")
	}

}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello world !")
}
