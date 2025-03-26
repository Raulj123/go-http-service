package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/Raulj123/go-service/config"
	"github.com/Raulj123/go-service/employee"
	"github.com/Raulj123/go-service/handler"
	"github.com/go-chi/chi"
	_ "github.com/mattn/go-sqlite3"
)

func initDB() *sql.DB{
	conf, err := config.Load("./env.json")
	if err != nil {
		fmt.Println("Could not load env",err)
	}
	db, err := sql.Open("sqlite3", conf.DBuri)
	if err != nil {
		log.Fatal(err)
	}
	err = db.Ping()
	if err != nil {
		log.Fatal("Could not verify connection to db")
	}
	return db
}

// Main loads config, opens db and runs http server
// TODO: add global logger
func main(){
	db := initDB()
	prov := employee.NewEmpProvider(db)
	r := chi.NewRouter()
	r.Mount("/", handler.NewHandler(prov))
	defer db.Close()
	log.Fatal("cannot serve:", http.ListenAndServe(":8080", r))
}