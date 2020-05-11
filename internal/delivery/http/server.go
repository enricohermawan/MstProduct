package http

import (
	"net/http"

	"product/pkg/grace"
)

// ProductHandler ...
type ProductHandler interface {
	// Masukkan fungsi handler di sini
	ProductHandler(w http.ResponseWriter, r *http.Request)
}

// Server ...
type Server struct {
	server  *http.Server
	Product ProductHandler
}

// Serve is serving HTTP gracefully on port x ...
func (s *Server) Serve(port string) error {
	return grace.Serve(port, s.Handler())
}
