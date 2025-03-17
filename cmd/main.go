package main

import (
	"log"
	"net/http"

	server "github.com/Raulj123/go-service/db"
	routes "github.com/Raulj123/go-service/internal/routes"
)

func main(){
	s := server.NewServer()
	defer s.Database.Close()
	routes.MountHandlers()
	log.Fatal(http.ListenAndServe(":3000", s.Router))
}

