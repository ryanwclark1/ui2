package server

import (
	"embed"
	"fmt"
	"io/fs"
	"log/slog"
	"net/http"

	"github.com/a-h/templ"

	"github.com/ryanwclark1/ui2/internal/middleware"
	"github.com/ryanwclark1/ui2/ui/pages"
)

func (s *Server) HandleAssets(assets embed.FS) http.Handler {
	contentAssets, err := fs.Sub(fs.FS(assets), "static")
	if err != nil {
		slog.Info("HandleAssets: failed to load assets: %v", err)
	}
	return http.StripPrefix("/static/", http.FileServerFS(contentAssets))
}

func (s *Server) HandleFavicon(assets embed.FS) http.Handler {
	b, err := assets.ReadFile("static/favicon.ico")
	if err != nil {
		slog.Info("HandleFavicon: failed to read favicon.ico: %v", err)
	}
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "image/x-icon")
		w.Write(b)
	})
}

func (s *Server) handlePageIndex() http.Handler {
	return templ.Handler(pages.DefaultHome, templ.WithContentType("text/html"))
}

func (s *Server) handleSaveSession() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		session, err := s.sess.Get(r, s.conf.CookieName)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		session.Values["foo"] = "bar"
		session.Values[42] = 43
		err = session.Save(r, w)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	})
}

func (s *Server) handleReadSession() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		session := middleware.SessionFromContext(r.Context())
		if session == nil {
			http.Error(w, "session not found", http.StatusInternalServerError)
			return
		}
		fmt.Fprint(w, session.Values)
	})
}
