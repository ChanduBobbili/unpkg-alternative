package main

import (
	"net/http"
	unpkghttp "unpkg-alternative/http"
	"unpkg-alternative/logs"

	"go.uber.org/zap"
)

func main() {
	logger := logs.NewLogger()
	defer logger.Sync()

	router := unpkghttp.NewRouter(logger)
	logger.Info("Server running", zap.String("url", "http://localhost:8080"))
	http.ListenAndServe(":8080", router)
}
