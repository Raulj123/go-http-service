package employee

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"

	"github.com/Raulj123/go-service/models/employee"
)

type mockHandler struct {}

func (m *mockHandler) Store(e employee.Employee) error {
	return nil
}

func (m *mockHandler) Employee(id int64) (*employee.Employee, error) {
	return &employee.Employee{
		Id: 90,
		Name: "TEST",
		Manager: "TEST MAN",
		StartDate: "testDate",
	}, nil
}

func (m *mockHandler) Employees() ([]employee.Employee, error) {
	return []employee.Employee{
		{Id: 001, Name: "TEST", Manager: "TEST MAN", StartDate: "testDate"},
	}, nil
}
func TestEmployeeHandler(t *testing.T) {
	expected := []employee.Employee{
		{Id: 1, Name: "TEST", Manager: "TEST MAN", StartDate: "testDate"},
	}
	req  := httptest.NewRequest("GET","/", nil)
	w := httptest.NewRecorder()
	h := employee.NewHandler(&mockHandler{})
	h.GetEmployees(w,req)
	
	if w.Code != http.StatusOK {
		t.Errorf("Expected status code %v, wanted %v", http.StatusOK, w.Code)
	}

	var actual []employee.Employee
	if err := json.Unmarshal(w.Body.Bytes(), &actual); err != nil {
		t.Fatalf("Failed to unmarshal response: %v", err)
	}

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("Expected response %v, got %v", expected, actual)
	}

}