package main

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/Raulj123/go-service/config"
	"github.com/Raulj123/go-service/models/employee"
	"github.com/go-chi/chi"
	_ "github.com/mattn/go-sqlite3"
)

func initDB() *sql.DB {
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
		log.Fatal("Could not verify connection to db",err)
	}
	return db
}

// Main loads config, opens db and runs http server
// TODO: add global logger
func main(){
	db := initDB()
	defer func() {
		if err := db.Close(); err != nil {
			log.Fatal("Could not close DB")
		}
	}()

	prov := employee.NewEmpProvider(db)
	r := chi.NewRouter()
	r.Mount("/", employee.NewHandler(prov))
	s := &http.Server{
		Addr: ":8080",
		Handler: r,
	}
	fmt.Println("starting server on localhost:8080")
	go func() {
		if err := s.ListenAndServe(); !errors.Is(err, http.ErrServerClosed) {
			log.Fatalf("HTTP server error: %v", err)
		}
		log.Println("no new connections")
	}()

	sigC := make(chan os.Signal, 1)
	signal.Notify(sigC, syscall.SIGINT, syscall.SIGTERM)
	<-sigC
	shutdownCtx, shutdownRelease := context.WithTimeout(context.Background(), 10*time.Second)
    defer shutdownRelease()

    if err := s.Shutdown(shutdownCtx); err != nil {
        log.Fatalf("HTTP shutdown error: %v", err)
    }
    log.Println("Graceful shutdown complete.")
}