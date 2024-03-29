package helpers

import (
	"log/slog"
	"net/http"
)

// StaticFileServerHandler handles a custom handler for serve embed static folder.
//
// Example:
//
//	import (
//		"embed"
//		"net/http"
//	)
//
//	//go:embed static/*
//	var static embed.FS
//
//	// Handle static files (with a custom handler).
//	http.Handle("/static/", helpers.StaticFileServerHandler(http.FS(static)))
func StaticFileServerHandler(fs http.FileSystem) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Check, if the requested file is existing.
		_, err := fs.Open(r.URL.Path)
		if err != nil {
			// If file is not found, return HTTP 404 error.
			http.NotFound(w, r)
			slog.Error(err.Error(), "method", r.Method, "status", http.StatusNotFound, "path", r.URL.Path)
			return
		}

		// File is found, return to standard http.FileServer.
		http.FileServer(fs).ServeHTTP(w, r)
	})
}