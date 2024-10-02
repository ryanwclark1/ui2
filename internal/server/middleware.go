package server

import (
	"net/http"


	"github.com/ryanwclark1/ui2/internal/middleware"
)

func (s *Server) mw(h http.Handler) http.Handler {
	middlewares := []func(http.Handler) http.Handler{
		middleware.CapturePath,
		middleware.CaptureHtmxRequestHeaders,
		middleware.Session(s.sess, s.conf.CookieName),
	}
	for _, m := range middlewares {
		h = m(h)
	}
	return h
}
