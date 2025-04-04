package http

import (
	"net/http"
	"unpkg-alternative/http/handler"
)

func NewRouter() http.Handler {
	// Create new HTTP request multiplexer
	mux := http.NewServeMux()
	// Register file handler at root
	mux.HandleFunc("/", handler.FileHandler)
	return mux
}
