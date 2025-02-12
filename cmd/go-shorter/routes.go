package main

import (
	"net/http"

	"github.com/Jalenarms1/go-shorter/internal/handlers"
)

func NewServer() *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("/new-url", handlers.HandleNewUrl)

	mux.HandleFunc("/go/{urlCode}", handlers.HandleRedirect)

	return mux
}
