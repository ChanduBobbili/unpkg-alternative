package handler

import (
	"net/http"
	"strings"
	"path/filepath"

	"unpkg-alternative/cache"
	"unpkg-alternative/npm"
)

// FileHandler processes incoming CDN file requests
func FileHandler(w http.ResponseWriter, r *http.Request) {
	// Extract "pkg@version/file.js" format
	parts := strings.SplitN(strings.TrimPrefix(r.URL.Path, "/"), "/", 2)
	if len(parts) != 2 {
		http.Error(w, "Invalid URL", http.StatusBadRequest)
		return
	}

	pkgVersion := parts[0]       // e.g., "react@18.2.0"
	filePath := parts[1]         // e.g., "index.js"
	name, version := npm.ParsePackage(pkgVersion)
	cachePath := filepath.Join("cache", pkgVersion) // Local cache path

	// Download & extract if not cached
	if !cache.Exists(cachePath) {
		err := npm.DownloadAndExtract(name, version, cachePath)
		if err != nil {
			http.Error(w, "Failed to fetch package", http.StatusInternalServerError)
			return
		}
	}

	// Serve the requested file
	http.ServeFile(w, r, filepath.Join(cachePath, "package", filePath))
}