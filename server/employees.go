package server

import (
	"fmt"
	"net/http"

	sqlite "github.com/Raulj123/go-service/sqlc"
	"github.com/Raulj123/go-service/utils"
)

// TODO context shi and queries blah blah oh and loogers?
// Employee represents an employee with minimal information
type Employee struct {
	Id int `json:"id"`
    Name string `json:"name"`
	Manager string `json:"manager"`
	StartDate string `json:"start_date"`
}
// This function handles getEmployees route in the future should handle all employee related routes
func (s *Server) getEmployees() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request){
		queries := sqlite.New(s.DB)
		employees, err := queries.GetAllEmployees(r.Context())
		defer r.Body.Close()
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
		queries := sqlite.New(s.DB)
		parmas := sqlite.GetEmployeesStartingSoonParams{
			StartDate: utils.GetCurrentDate(),
			StartDate_2: utils.GetFutureDate(),
		}
		employees, err := queries.GetEmployeesStartingSoon(r.Context(), parmas)
		defer r.Body.Close()
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

func (s *Server) postEmployee() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request){
		if r.Method != http.MethodPost {
			http.Error(w, "Invalid method",http.StatusBadRequest)
			return
		}
		
		d,err := utils.DecodeJson[Employee](r)
		
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		queries := sqlite.New(s.DB)
		params := sqlite.InsertEmployeeParams{
			ID: int64(d.Id),
			Name: d.Name,
			Manager: d.Manager,
			StartDate: d.StartDate,
		}
		employee, err := queries.InsertEmployee(r.Context(), params)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		fmt.Println("Employee: ", employee)
		// does go infer the type here?
		utils.EncodeJson(w,http.StatusOK, d)
	}
}