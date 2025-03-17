package db

import (
	"database/sql"
	"log"

	"github.com/go-chi/chi/v5"
	_ "github.com/mattn/go-sqlite3"
)

type Server struct {
	Router *chi.Mux
	Database *sql.DB
}

var globalServer *Server

// Sets up router and db
func NewServer() *Server {
	globalServer = &Server{}
	globalServer.Router = chi.NewRouter()
	db, err := sql.Open("sqlite3", "./sqlite.db")
	if err != nil {
		log.Fatal(err)
	}
	err = db.Ping()
	if err != nil {
		log.Fatal("Could not verify connection to db")
	}
	globalServer.Database = db
	return globalServer
}

// Returns var pointing to struct
func GetConn() *Server {
	return globalServer
}