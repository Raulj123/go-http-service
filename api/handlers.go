package api

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/Raulj123/go-service/db"
	sqlite "github.com/Raulj123/go-service/internal/sqlite"
)

func GetEmployees(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	db := db.GetConn()
	queries := sqlite.New(db.Database)
	employees, err := queries.GetAllEmployees(ctx)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	for _ , e := range employees {
		fmt.Println("Employee starting soon", "Name", e.Name, "StartDate", e.StartDate)
	}
	w.Header().Set("Content-Type", "application/json")
    
    if err := json.NewEncoder(w).Encode(employees); err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
}
