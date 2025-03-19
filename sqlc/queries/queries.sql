-- name: InsertEmployee :one
INSERT INTO employees (id, name, manager, start_date)
VALUES (?, ?, ?, ?)
RETURNING *;

-- name: GetAllEmployees :many
SELECT * FROM employees;

-- name: GetEmployeesStartingSoon :many
SELECT * FROM employees WHERE start_date >= ? AND start_date <= ?;