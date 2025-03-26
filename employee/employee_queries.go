package employee

import (
	"context"
	"database/sql"

	sqlite "github.com/Raulj123/go-service/sqlc"
)

type empProvider struct {
	*sql.DB
	*sqlite.Queries
}

var ctx = context.Background()

func NewEmpProvider(db *sql.DB) EmpProvider {
	queries := sqlite.New(db)
	return empProvider{db, queries}
}

func (e empProvider) Employee(id int64) (*Employee, error) {
	emp, err := e.Queries.GetEmployee(ctx, id)
	if err != nil {
		return nil,err
	}
	return &Employee{
		Id: int(emp.ID),
		Name: emp.Name,
		Manager: emp.Manager,
		StartDate: emp.StartDate,
	}, nil
}

func (e empProvider) Store(emp Employee) error {
	_, err := e.Queries.InsertEmployee(ctx, sqlite.InsertEmployeeParams{
		ID: int64(emp.Id),
		Name: emp.Name,
		Manager: emp.Manager,
		StartDate: emp.StartDate,
	})
	if err != nil {
		return err
	}
	return nil
}
