package server

import (
	"database/sql"
	"log"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	_ "github.com/mattn/go-sqlite3"
)

type Server struct {
	server *http.Server
	Router *chi.Mux
	DB     *sql.DB
}

// This function creates new server struct, calls routes and returns server to main
func NewServer(db *sql.DB) *Server {
	s := Server {
		server: &http.Server{
			WriteTimeout: 5 *time.Second,
			ReadTimeout: 5 * time.Second,
			IdleTimeout: 5 *time.Second,
		},
		Router: chi.NewRouter(),
	}
	s.routes()
	s.server.Handler = s.Router
	s.DB = db
	return &s
}

// This function runs the server
func (s *Server) Run(port string) error {
	s.server.Addr = port
	log.Printf("Server starting on %s", port)
	return s.server.ListenAndServe()
}
