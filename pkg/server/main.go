package server

import (
	"net/http"

	"github.com/gorilla/mux"
)

// Server struct
type Server struct {
	r *mux.Router
}

// NewServer returns Server struct
func NewServer() *Server {
	s := Server{r: mux.NewRouter()}
	s.NewRouter()
	return &s
}

// Start starts server
func (s *Server) Start() {
	if err := http.ListenAndServe(":8080", s.r); err != nil {
		// handle error
	}
}
