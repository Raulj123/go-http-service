package employee

type Employee struct {
	Id int `json:"id"`
    Name string `json:"name"`
	Manager string `json:"manager"`
	StartDate string `json:"start_date"`
}

type Provider interface {
	Store(Employee) error 
	Employee(id int64) (*Employee, error)
	Employees() ([]Employee, error)
}