package http

import (
	"net/http"
	"unpkg-alternative/http/handler"

	"go.uber.org/zap"
)

func NewRouter(logger *zap.Logger) http.Handler {
	mux := http.NewServeMux()

	// Logging middleware
	loggedHandler := func(h http.HandlerFunc) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			logger.Info("Request", zap.String("url", r.URL.Path))
			h(w, r)
		}
	}

	mux.HandleFunc("/", loggedHandler(handler.FileHandler(logger)))
	return mux
}
