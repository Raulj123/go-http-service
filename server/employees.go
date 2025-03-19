package server

import (
	"context"
	"fmt"
	"net/http"

	sqlite "github.com/Raulj123/go-service/sqlc"
)

// This function handles getEmployees route in the future should handle all employee related routes
func (s *Server) getEmployees() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request){
		ctx := context.Background()
		queries := sqlite.New(s.DB)
		employees, err := queries.GetAllEmployees(ctx)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		// TODO: add validator, encoder/decoder func 
		for _, e := range employees {
			fmt.Fprintf(w, "Name %s, starts: %s", e.Name, e.StartDate.String)
		}
	}
}