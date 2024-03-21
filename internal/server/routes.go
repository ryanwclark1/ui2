package server

import (
	"embed"
)


//go:embed static
var static embed.FS

func (s *Server) Routes() error {
	s.r.Handle("GET /static/*", s.mw(s.HandleAssets(static)))
	s.r.Handle("GET /favicon.ico", s.HandleFavicon(static))

	s.r.Handle("GET /session", s.mw(s.handleSaveSession()))
	s.r.Handle("GET /read-session", s.mw(s.handleReadSession()))

	s.r.Handle("GET /", s.mw(s.handlePageIndex()))

	s.r.Handle("GET /form", s.mw(s.handlePageForm()))
	s.r.Handle("GET /form/submit", s.mw(s.handleFormSubmit()))
	s.r.Handle("POST /form/submit", s.mw(s.handleFormSubmit()))

	s.r.Handle("GET /search", s.mw(s.handlePageSearch()))
	s.r.Handle("POST /search/users", s.mw(s.handleSearchUsers()))

	return nil
}
