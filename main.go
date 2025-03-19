package main

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/Raulj123/go-service/config"
	"github.com/Raulj123/go-service/server"
	_ "github.com/mattn/go-sqlite3"
)

// Main loads config, opens db and calls NewServer and runs http server
// TODO: add global logger
func main(){
	conf, err := config.LoadConfig("./env.json")
	if err != nil {
		fmt.Println("Could not load env",err)
	}
	db, err := sql.Open("sqlite3", conf.DBuri)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	err = db.Ping()
	if err != nil {
		log.Fatal("Could not verify connection to db")
	}
	srv := server.NewServer(db)
	log.Fatal(srv.Run(conf.Port))
}