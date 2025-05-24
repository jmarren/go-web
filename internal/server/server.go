package server

import "net/http"

type Server struct {
	*http.Server
}

func New() *Server {
	s := &Server{
		&http.Server{
			Addr:    ":8080",
			Handler: &http.ServeMux{},
		},
	}
	return s
}

func (s *Server) Go() {
	s.ListenAndServe()
}
