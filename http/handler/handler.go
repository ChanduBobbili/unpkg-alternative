package handler

import (
	"net/http"
	"path/filepath"
	"strings"

	"unpkg-alternative/cache"
	"unpkg-alternative/npm"

	"go.uber.org/zap"
)

func FileHandler(logger *zap.Logger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		parts := strings.SplitN(strings.TrimPrefix(r.URL.Path, "/"), "/", 2)
		if len(parts) != 2 {
			http.Error(w, "Invalid URL", http.StatusBadRequest)
			logger.Warn("Malformed URL", zap.String("url", r.URL.Path))
			return
		}

		pkgVersion := parts[0]
		filePath := parts[1]
		name, version := npm.ParsePackage(pkgVersion)
		cachePath := filepath.Join("cache", pkgVersion)

		if !cache.Exists(cachePath) {
			err := npm.DownloadAndExtract(name, version, cachePath)
			if err != nil {
				http.Error(w, "Failed to fetch package", http.StatusInternalServerError)
				logger.Error("Download error", zap.Error(err))
				return
			}
		}

		fullFilePath := filepath.Join(cachePath, "package", filePath)
		http.ServeFile(w, r, fullFilePath)
		logger.Info("Served file", zap.String("path", fullFilePath))
	}
}
