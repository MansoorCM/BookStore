package api

import (
	"fmt"
	"net/http"
)

func OrderHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Here's your order summary.")
}
