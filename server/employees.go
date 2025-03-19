package server

import (
	"context"
	"fmt"
	"net/http"

	sqlite "github.com/Raulj123/go-service/sqlc"
	"github.com/Raulj123/go-service/utils"
)

// TODO context shi and queries blah blah oh and loogers?

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
			fmt.Fprintf(w, "Name %s, starts: %s", e.Name, e.StartDate)
		}
	}
}

func (s *Server) getStartSoon() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request){
		ctx := context.Background()
		queries := sqlite.New(s.DB)
		parmas := sqlite.GetEmployeesStartingSoonParams{
			StartDate: utils.GetCurrentDate(),
			StartDate_2: utils.GetFutureDate(),
		}
		employees, err := queries.GetEmployeesStartingSoon(ctx, parmas)
		
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		// TODO: add validator, encoder/decoder func
		if len(employees) != 0 { 
		for _, e := range employees {
			fmt.Fprintf(w, "Name %s, starts: %s", e.Name, e.StartDate)
		}
	} else {
		fmt.Fprintf(w, "No Employees starting soon")
	}
	}
}