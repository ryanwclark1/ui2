package server

import (
	"net/http"

	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"

	m "github.com/ryanwclark1/ui2/internal/middleware"
)

func (s *Server) mw(h http.Handler) http.Handler {
	middlewares := []func(http.Handler) http.Handler{
		cors.Handler(cors.Options{
			AllowedOrigins:   s.conf.AllowedOrigins,
			AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
			AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
			ExposedHeaders:   []string{"Link"},
			AllowCredentials: true,
			MaxAge:           300,
		}),
		middleware.RedirectSlashes,
		middleware.Recoverer,
		middleware.Compress(5),
		middleware.Logger,
		m.CapturePath,
		m.CaptureHtmxRequestHeaders,
		m.Session(s.sess, s.conf.CookieName),
	}
	for _, m := range middlewares {
		h = m(h)
	}
	return h
}
