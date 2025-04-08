package api

import (
	"fmt"
	"net/http"
)

type MessageHandler struct {
	Message string
}

func (m *MessageHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, m.Message)
}
