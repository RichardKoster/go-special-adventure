package web

import (
	"github.com/go-chi/chi"
	"net/http"
)

func (s *Server) mountRoutes(r chi.Router) {
	r.Get("/", healthStatus)
}

func healthStatus(w http.ResponseWriter, r *http.Request) {
	writeString(w, http.StatusOK, "Hello World");
}