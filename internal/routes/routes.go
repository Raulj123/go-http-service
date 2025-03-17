package routes

import (
	"github.com/Raulj123/go-service/api"
	"github.com/Raulj123/go-service/db"
)

func MountHandlers() {
	server := db.GetConn()
	server.Router.Get("/employees", api.GetEmployees)
	// s.Router.Get("/employee/{id}",api.GetUser)
}