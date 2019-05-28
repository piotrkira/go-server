package server

import (
	"net/http"
)

// NewRouter holds possible url addreses
func (s *Server) NewRouter() {
	// Web
	s.r.HandleFunc("/", s.indexHandler())

	// Static
	s.r.PathPrefix("/assets/").Handler(http.StripPrefix("/assets/", http.FileServer(http.Dir("path/to/local/directory"))))
}
